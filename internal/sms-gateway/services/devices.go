package services

import (
	"github.com/capcom6/sms-gateway/internal/sms-gateway/models"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/repositories"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type DevicesServiceParams struct {
	fx.In

	Devices *repositories.DevicesRepository
	Logger  *zap.Logger
}

type DevicesService struct {
	Devices *repositories.DevicesRepository

	Logger *zap.Logger
}

func (s *DevicesService) Select(user models.User) ([]models.Device, error) {
	return s.Devices.Select(repositories.SelectDevicesFilter{UserId: &user.ID})
}

func NewDevicesService(params DevicesServiceParams) *DevicesService {
	return &DevicesService{
		Devices: params.Devices,
		Logger:  params.Logger.Named("DevicesService"),
	}
}
