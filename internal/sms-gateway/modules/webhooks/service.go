package webhooks

import (
	"fmt"

	"github.com/android-sms-gateway/client-go/smsgateway/webhooks"
	"github.com/capcom6/go-helpers/slices"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/modules/db"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/modules/devices"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/modules/push"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type ServiceParams struct {
	fx.In

	IDGen db.IDGen

	Webhooks *Repository

	DevicesSvc *devices.Service
	PushSvc    *push.Service

	Logger *zap.Logger
}

type Service struct {
	idgen db.IDGen

	webhooks *Repository

	devicesSvc *devices.Service
	pushSvc    *push.Service

	logger *zap.Logger
}

func NewService(params ServiceParams) *Service {
	return &Service{
		idgen:      params.IDGen,
		webhooks:   params.Webhooks,
		devicesSvc: params.DevicesSvc,
		pushSvc:    params.PushSvc,
		logger:     params.Logger,
	}
}

func (s *Service) Select(userID string, filters ...SelectFilter) ([]webhooks.Webhook, error) {
	filters = append(filters, WithUserID(userID))

	items, err := s.webhooks.Select(filters...)
	if err != nil {
		return nil, fmt.Errorf("can't select webhooks: %w", err)
	}

	return slices.Map(items, webhookToDTO), nil
}

func (s *Service) Replace(userID string, webhook *webhooks.Webhook) error {
	if !webhooks.IsValidEventType(webhook.Event) {
		return newValidationError("event", string(webhook.Event), fmt.Errorf("enum value expected"))
	}

	if webhook.ID == "" {
		webhook.ID = s.idgen()
	}

	model := Webhook{
		ExtID:  webhook.ID,
		UserID: userID,
		URL:    webhook.URL,
		Event:  webhook.Event,
	}

	if err := s.webhooks.Replace(&model); err != nil {
		return fmt.Errorf("can't replace webhook: %w", err)
	}

	go s.notifyDevices(userID)

	return nil
}

func (s *Service) Delete(userID string, filters ...SelectFilter) error {
	filters = append(filters, WithUserID(userID))
	if err := s.webhooks.Delete(filters...); err != nil {
		return fmt.Errorf("can't delete webhooks: %w", err)
	}

	go s.notifyDevices(userID)

	return nil
}

func (s *Service) notifyDevices(userID string) {
	s.logger.Info("Notifying devices", zap.String("user_id", userID))

	devices, err := s.devicesSvc.Select(devices.WithUserID(userID))
	if err != nil {
		s.logger.Error("Failed to select devices", zap.String("user_id", userID), zap.Error(err))
		return
	}

	for _, device := range devices {
		if device.PushToken == nil {
			continue
		}

		if err := s.pushSvc.Enqueue(*device.PushToken, push.NewWebhooksUpdatedEvent()); err != nil {
			s.logger.Error("Failed to send push notification", zap.String("user_id", userID), zap.Error(err))
		}
	}

	s.logger.Info("Notified devices", zap.String("user_id", userID), zap.Int("count", len(devices)))
}
