package services

import (
	"bitbucket.org/capcom6/smsgatewaybackend/internal/smsgateway/models"
	"bitbucket.org/capcom6/smsgatewaybackend/internal/smsgateway/repositories"
	"bitbucket.org/capcom6/smsgatewaybackend/pkg/smsgateway"
	"github.com/jaevor/go-nanoid"
)

type MessagesService struct {
	messages *repositories.MessagesRepository

	idgen func() string
}

func (s *MessagesService) SelectPending(deviceID string) ([]smsgateway.Message, error) {
	messages, err := s.messages.SelectPending(deviceID)
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
	existing, err := s.messages.Get(deviceID, message.ID)
	if err != nil {
		return err
	}

	existing.State = models.MessageState(message.State)
	existing.Recipients = s.recipientsStateToModel(message.Recipients)

	return s.messages.UpdateState(&existing)
}

func (s *MessagesService) Enqeue(deviceID string, message smsgateway.Message) error {
	msg := models.Message{
		DeviceID:   deviceID,
		ExtID:      message.ID,
		Message:    message.Message,
		Recipients: s.recipientsToModel(message.PhoneNumbers),
	}
	if msg.ExtID == "" {
		msg.ExtID = s.idgen()
	}

	return s.messages.Insert(&msg)
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

func NewMessagesService(messages *repositories.MessagesRepository) *MessagesService {
	idgen, _ := nanoid.Standard(21)

	return &MessagesService{
		messages: messages,
		idgen:    idgen,
	}
}
