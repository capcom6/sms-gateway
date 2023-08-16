package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"bitbucket.org/capcom6/smsgatewaybackend/internal/smsgateway/models"
	"bitbucket.org/capcom6/smsgatewaybackend/internal/smsgateway/repositories"
	"bitbucket.org/capcom6/smsgatewaybackend/pkg/smsgateway"
	"bitbucket.org/soft-c/gohelpers/pkg/filters"
	"github.com/jaevor/go-nanoid"
)

type MessagesService struct {
	Messages *repositories.MessagesRepository
	PushSvc  *PushService

	idgen func() string
}

func NewMessagesService(pushSvc *PushService, messages *repositories.MessagesRepository) *MessagesService {
	idgen, _ := nanoid.Standard(21)

	return &MessagesService{
		Messages: messages,
		PushSvc:  pushSvc,
		idgen:    idgen,
	}
}

func (s *MessagesService) SelectPending(deviceID string) ([]smsgateway.Message, error) {
	messages, err := s.Messages.SelectPending(deviceID)
	if err != nil {
		return nil, err
	}

	result := make([]smsgateway.Message, len(messages))
	for i, v := range messages {
		result[i] = smsgateway.Message{
			ID:           v.ExtID,
			Message:      v.Message,
			PhoneNumbers: s.recipientsToDomain(v.Recipients),
		}
	}

	return result, nil
}

func (s *MessagesService) UpdateState(deviceID string, message smsgateway.MessageState) error {
	existing, err := s.Messages.Get(deviceID, message.ID)
	if err != nil {
		return err
	}

	existing.State = models.MessageState(message.State)
	existing.Recipients = s.recipientsStateToModel(message.Recipients)

	return s.Messages.UpdateState(&existing)
}

func (s *MessagesService) Enqeue(device models.Device, message smsgateway.Message) error {
	for i, v := range message.PhoneNumbers {
		phone, err := filters.FilterPhone(v, false)
		if err != nil {
			return fmt.Errorf("некорректный номер телефона в строке %d: %w", i+1, err)
		}
		message.PhoneNumbers[i] = phone
	}

	msg := models.Message{
		DeviceID:   device.ID,
		ExtID:      message.ID,
		Message:    message.Message,
		Recipients: s.recipientsToModel(message.PhoneNumbers),
	}
	if msg.ExtID == "" {
		msg.ExtID = s.idgen()
	}

	if err := s.Messages.Insert(&msg); err != nil {
		return err
	}

	if device.PushToken == nil {
		return nil
	}

	go func(token string) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := s.PushSvc.Send(ctx, token, map[string]string{}); err != nil {
			log.Printf("failed to send push to %s: %v", *device.PushToken, err)
		}
	}(*device.PushToken)

	return nil
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

func (s *MessagesService) recipientsStateToModel(input []smsgateway.RecipientState) []models.MessageRecipient {
	output := make([]models.MessageRecipient, len(input))

	for i, v := range input {
		output[i] = models.MessageRecipient{
			PhoneNumber: v.PhoneNumber,
			State:       models.MessageState(v.State),
		}
	}

	return output
}
