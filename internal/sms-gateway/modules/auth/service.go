package auth

import (
	"fmt"

	"github.com/capcom6/sms-gateway/internal/sms-gateway/models"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/modules/devices"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/repositories"
	"github.com/capcom6/sms-gateway/pkg/crypto"
	"github.com/jaevor/go-nanoid"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Config struct {
	Mode         Mode
	PrivateToken string
}

type Params struct {
	fx.In

	Config Config

	Users      *repositories.UsersRepository
	DevicesSvc *devices.Service

	Logger *zap.Logger
}

type Service struct {
	config Config

	users      *repositories.UsersRepository
	devicesSvc *devices.Service

	logger *zap.Logger

	idgen func() string
}

func New(params Params) *Service {
	idgen, _ := nanoid.Standard(21)

	return &Service{
		config:     params.Config,
		users:      params.Users,
		devicesSvc: params.DevicesSvc,
		logger:     params.Logger.Named("Service"),
		idgen:      idgen,
	}
}

func (s *Service) RegisterUser(login, password string) (models.User, error) {
	user := models.User{
		ID: login,
	}

	var err error
	if user.PasswordHash, err = crypto.MakeBCryptHash(password); err != nil {
		return user, fmt.Errorf("can't hash password: %w", err)
	}

	if err = s.users.Insert(&user); err != nil {
		return user, fmt.Errorf("can't create user")
	}

	return user, nil
}

func (s *Service) RegisterDevice(user models.User, name, pushToken *string) (models.Device, error) {
	device := models.Device{
		Name:      name,
		PushToken: pushToken,
	}

	return device, s.devicesSvc.Insert(user, &device)
}

func (s *Service) UpdateDevice(id, pushToken string) error {
	return s.devicesSvc.UpdateToken(id, pushToken)
}

func (s *Service) IsPublic() bool {
	return s.config.Mode == ModePublic
}

func (s *Service) AuthorizeRegistration(token string) error {
	if s.IsPublic() {
		return nil
	}

	if token == s.config.PrivateToken {
		return nil
	}

	return fmt.Errorf("invalid token")
}

func (s *Service) AuthorizeDevice(token string) (models.Device, error) {
	device, err := s.devicesSvc.Get(devices.WithToken(token))
	if err != nil {
		return device, err
	}

	if err := s.devicesSvc.UpdateLastSeen(device.ID); err != nil {
		s.logger.Error("can't update last seen", zap.Error(err))
	}

	return device, nil
}

func (s *Service) AuthorizeUser(username, password string) (models.User, error) {
	user, err := s.users.GetByLogin(username)
	if err != nil {
		return user, err
	}

	return user, crypto.CompareBCryptHash(user.PasswordHash, password)
}
