package webhooks

import (
	"fmt"

	"github.com/capcom6/go-helpers/slices"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/modules/db"
)

type Service struct {
	idgen db.IDGen

	webhooks *Repository
}

func NewService(idgen db.IDGen, webhooks *Repository) *Service {
	return &Service{
		idgen:    idgen,
		webhooks: webhooks,
	}
}

func (s *Service) Select(userID string, filters ...SelectFilter) ([]WebhookDTO, error) {
	filters = append(filters, WithUserID(userID))

	items, err := s.webhooks.Select(filters...)
	if err != nil {
		return nil, fmt.Errorf("can't select webhooks: %w", err)
	}

	return slices.Map(items, webhookToDTO), nil
}

func (s *Service) Replace(userID string, webhook WebhookDTO) error {
	if webhook.ID == "" {
		webhook.ID = s.idgen()
	}

	model := Webhook{
		ExtID:  webhook.ID,
		UserID: userID,
		URL:    webhook.URL,
		Event:  webhook.Event,
	}

	return s.webhooks.Replace(&model)
}

func (s *Service) Delete(userID string, filters ...SelectFilter) error {
	filters = append(filters, WithUserID(userID))
	return s.webhooks.Delete(filters...)
}
