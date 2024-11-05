package devices

import (
	"context"
	"time"

	"github.com/capcom6/sms-gateway/internal/sms-gateway/models"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/modules/db"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type ServiceParams struct {
	fx.In

	Config Config

	Devices *repository

	IDGen db.IDGen

	Logger *zap.Logger
}

type Service struct {
	config Config

	devices *repository

	idGen db.IDGen

	logger *zap.Logger
}

func (s *Service) Insert(user models.User, device *models.Device) error {
	device.ID = s.idGen()
	device.AuthToken = s.idGen()
	device.UserID = user.ID

	return s.devices.Insert(device)
}

func (s *Service) Select(filter ...SelectFilter) ([]models.Device, error) {
	return s.devices.Select(filter...)
}

func (s *Service) Get(filter ...SelectFilter) (models.Device, error) {
	return s.devices.Get(filter...)
}

func (s *Service) UpdateToken(deviceId string, token string) error {
	return s.devices.UpdateToken(deviceId, token)
}

func (s *Service) UpdateLastSeen(deviceId string) error {
	return s.devices.UpdateLastSeen(deviceId)
}

func (s *Service) Clean(ctx context.Context) error {
	n, err := s.devices.removeUnused(ctx, time.Now().Add(-s.config.UnusedLifetime))

	s.logger.Info("Cleaned unused devices", zap.Int64("count", n))
	return err
}

func NewService(params ServiceParams) *Service {
	return &Service{
		config:  params.Config,
		devices: params.Devices,
		idGen:   params.IDGen,
		logger:  params.Logger.Named("service"),
	}
}
