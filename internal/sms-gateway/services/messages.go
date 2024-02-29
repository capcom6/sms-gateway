package services

import (
	"context"
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/capcom6/sms-gateway/internal/sms-gateway/models"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/repositories"
	"github.com/capcom6/sms-gateway/pkg/slices"
	"github.com/capcom6/sms-gateway/pkg/smsgateway"
	"github.com/capcom6/sms-gateway/pkg/types"
	"github.com/jaevor/go-nanoid"
	"github.com/nyaruka/phonenumbers"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

const (
	ErrorTTLExpired = "TTL expired"
)

type ErrValidation string

func (e ErrValidation) Error() string {
	return string(e)
}

type MessagesEnqueueOptions struct {
	SkipPhoneValidation bool
}

type MessagesServiceParams struct {
	fx.In

	Messages *repositories.MessagesRepository
	PushSvc  *PushService
	Logger   *zap.Logger
}

type MessagesService struct {
	Messages *repositories.MessagesRepository
	PushSvc  *PushService
	Logger   *zap.Logger

	idgen func() string
}

func NewMessagesService(params MessagesServiceParams) *MessagesService {
	idgen, _ := nanoid.Standard(21)

	return &MessagesService{
		Messages: params.Messages,
		PushSvc:  params.PushSvc,
		Logger:   params.Logger.Named("MessagesService"),
		idgen:    idgen,
	}
}

func (s *MessagesService) SelectPending(deviceID string) ([]smsgateway.Message, error) {
	messages, err := s.Messages.SelectPending(deviceID)
	if err != nil {
		return nil, err
	}

	messages = s.filterTimeouted(messages)

	result := make([]smsgateway.Message, len(messages))
	for i, v := range messages {
		var ttl *uint64 = nil
		if v.ValidUntil != nil {
			delta := time.Until(*v.ValidUntil).Seconds()
			if delta > 0 {
				deltaInt := uint64(delta)
				ttl = &deltaInt
			} else {
				deltaInt := uint64(0)
				ttl = &deltaInt
			}
		}

		result[i] = smsgateway.Message{
			ID:                 v.ExtID,
			Message:            v.Message,
			SimNumber:          v.SimNumber,
			WithDeliveryReport: types.AsPointer[bool](v.WithDeliveryReport),
			IsEncrypted:        v.IsEncrypted,
			PhoneNumbers:       s.recipientsToDomain(v.Recipients),
			TTL:                ttl,
			ValidUntil:         v.ValidUntil,
		}
	}

	return result, nil
}

func (s *MessagesService) UpdateState(deviceID string, message smsgateway.MessageState) error {
	existing, err := s.Messages.Get(message.ID, repositories.MessagesSelectFilter{DeviceID: deviceID})
	if err != nil {
		return err
	}

	if message.State == smsgateway.MessageStatePending {
		message.State = smsgateway.MessageStateProcessed
	}

	existing.State = models.MessageState(message.State)
	existing.Recipients = s.recipientsStateToModel(message.Recipients, existing.IsHashed)

	return s.Messages.UpdateState(&existing)
}

func (s *MessagesService) GetState(user models.User, ID string) (smsgateway.MessageState, error) {
	message, err := s.Messages.Get(ID, repositories.MessagesSelectFilter{}, repositories.MessagesSelectOptions{WithRecipients: true, WithDevice: true})
	if err != nil {
		return smsgateway.MessageState{}, repositories.ErrMessageNotFound
	}

	if message.Device.UserID != user.ID {
		return smsgateway.MessageState{}, repositories.ErrMessageNotFound
	}

	return modelToMessageState(message), nil
}

func (s *MessagesService) Enqeue(device models.Device, message smsgateway.Message, opts MessagesEnqueueOptions) (smsgateway.MessageState, error) {
	state := smsgateway.MessageState{
		ID:         "",
		State:      smsgateway.MessageStatePending,
		Recipients: make([]smsgateway.RecipientState, len(message.PhoneNumbers)),
	}

	var phone string
	var err error
	for i, v := range message.PhoneNumbers {
		if message.IsEncrypted || opts.SkipPhoneValidation {
			phone = v
		} else {
			if phone, err = cleanPhoneNumber(v); err != nil {
				return state, fmt.Errorf("can't use phone in row %d: %w", i+1, err)
			}
		}

		message.PhoneNumbers[i] = phone

		state.Recipients[i] = smsgateway.RecipientState{
			PhoneNumber: phone,
			State:       smsgateway.MessageStatePending,
		}
	}

	var validUntil *time.Time = message.ValidUntil
	if message.TTL != nil && *message.TTL > 0 {
		validUntil = types.AsPointer(time.Now().Add(time.Duration(*message.TTL) * time.Second))
	}

	msg := models.Message{
		DeviceID:           device.ID,
		ExtID:              message.ID,
		Message:            message.Message,
		ValidUntil:         validUntil,
		SimNumber:          message.SimNumber,
		WithDeliveryReport: types.OrDefault[bool](message.WithDeliveryReport, true),
		IsEncrypted:        message.IsEncrypted,
		Device:             device,
		Recipients:         s.recipientsToModel(message.PhoneNumbers),
		TimedModel:         models.TimedModel{},
	}
	if msg.ExtID == "" {
		msg.ExtID = s.idgen()
	}
	state.ID = msg.ExtID

	if err := s.Messages.Insert(&msg); err != nil {
		return state, err
	}

	if device.PushToken == nil {
		return state, nil
	}

	go func(token string) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := s.PushSvc.Enqueue(ctx, token, map[string]string{}); err != nil {
			s.Logger.Error("Can't enqueue message", zap.String("token", token), zap.Error(err))
		}
	}(*device.PushToken)

	return state, nil
}

func (s *MessagesService) HashProcessed() error {
	return s.Messages.HashProcessed()
}

func (s *MessagesService) filterTimeouted(messages []models.Message) []models.Message {
	result := make([]models.Message, 0, len(messages))
	for _, v := range messages {
		if v.ValidUntil == nil || time.Now().Before(*v.ValidUntil) {
			result = append(result, v)
		} else if v.State == models.MessageStatePending {
			v.State = models.MessageStateFailed
			for i := range v.Recipients {
				v.Recipients[i].State = models.MessageStateFailed
				v.Recipients[i].Error = types.AsPointer(ErrorTTLExpired)
			}
			if err := s.Messages.UpdateState(&v); err != nil {
				s.Logger.Error("Can't update message state", zap.Error(err))
			}
		}
	}
	return result
}

func (s *MessagesService) recipientsToDomain(input []models.MessageRecipient) []string {
	output := make([]string, len(input))

	for i, v := range input {
		output[i] = v.PhoneNumber
	}

	return output
}

func (s *MessagesService) recipientsToModel(input []string) []models.MessageRecipient {
	output := make([]models.MessageRecipient, len(input))

	for i, v := range input {
		output[i] = models.MessageRecipient{
			PhoneNumber: v,
		}
	}

	return output
}

func (s *MessagesService) recipientsStateToModel(input []smsgateway.RecipientState, hash bool) []models.MessageRecipient {
	output := make([]models.MessageRecipient, len(input))

	for i, v := range input {
		phoneNumber := v.PhoneNumber
		if len(phoneNumber) > 0 && phoneNumber[0] != '+' {
			// compatibility with Android app before 1.1.1
			phoneNumber = "+" + phoneNumber
		}

		if v.State == smsgateway.MessageStatePending {
			v.State = smsgateway.MessageStateProcessed
		}

		if hash {
			phoneNumber = fmt.Sprintf("%x", sha256.Sum256([]byte(phoneNumber)))[:16]
		}

		output[i] = models.MessageRecipient{
			PhoneNumber: phoneNumber,
			State:       models.MessageState(v.State),
			Error:       v.Error,
		}
	}

	return output
}

func modelToMessageState(input models.Message) smsgateway.MessageState {
	return smsgateway.MessageState{
		ID:          input.ExtID,
		State:       smsgateway.ProcessState(input.State),
		IsHashed:    input.IsHashed,
		IsEncrypted: input.IsEncrypted,
		Recipients:  slices.Map(input.Recipients, modelToRecipientState),
	}
}

func modelToRecipientState(input models.MessageRecipient) smsgateway.RecipientState {
	return smsgateway.RecipientState{
		PhoneNumber: input.PhoneNumber,
		State:       smsgateway.ProcessState(input.State),
		Error:       input.Error,
	}
}

func cleanPhoneNumber(input string) (string, error) {
	phone, err := phonenumbers.Parse(input, "RU")
	if err != nil {
		return input, ErrValidation(fmt.Sprintf("can't parse phone number: %s", err.Error()))
	}

	if !phonenumbers.IsValidNumber(phone) {
		return input, ErrValidation("invalid phone number")
	}

	phoneNumberType := phonenumbers.GetNumberType(phone)
	if phoneNumberType != phonenumbers.MOBILE && phoneNumberType != phonenumbers.FIXED_LINE_OR_MOBILE {
		return input, ErrValidation("not mobile phone number")
	}

	return phonenumbers.Format(phone, phonenumbers.E164), nil
}
