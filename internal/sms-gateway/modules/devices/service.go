package devices

import (
	"github.com/capcom6/sms-gateway/internal/sms-gateway/models"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/modules/db"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type ServiceParams struct {
	fx.In

	Devices *repository

	IDGen db.IDGen

	Logger *zap.Logger
}

type Service struct {
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

func NewService(params ServiceParams) *Service {
	return &Service{
		devices: params.Devices,
		idGen:   params.IDGen,
		logger:  params.Logger.Named("service"),
	}
}
