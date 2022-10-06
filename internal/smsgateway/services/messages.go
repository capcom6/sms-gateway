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

func (s *MessagesService) recipientsToDomain(input []models.MessageRecipient) []string {
	output := make([]string, len(input))

	for i, v := range input {
		output[i] = v.PhoneNumber
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
