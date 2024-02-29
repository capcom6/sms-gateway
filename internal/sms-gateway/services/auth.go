package services

import (
	"fmt"

	"github.com/capcom6/sms-gateway/internal/sms-gateway/models"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/repositories"
	"github.com/capcom6/sms-gateway/pkg/crypto"
	"github.com/jaevor/go-nanoid"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type AuthService struct {
	users   *repositories.UsersRepository
	devices *repositories.DevicesRepository

	logger *zap.Logger

	idgen func() string
}

func (s *AuthService) RegisterUser(login, password string) (models.User, error) {
	user := models.User{
		ID: login,
	}

	var err error
	if user.PasswordHash, err = crypto.MakeBCryptHash(password); err != nil {
		return user, err
	}

	if err = s.users.Insert(&user); err != nil {
		return user, fmt.Errorf("can't create user")
	}

	return user, nil
}

func (s *AuthService) RegisterDevice(userID string, name, pushToken *string) (models.Device, error) {
	device := models.Device{
		ID:        s.idgen(),
		Name:      name,
		AuthToken: s.idgen(),
		PushToken: pushToken,
		UserID:    userID,
	}

	return device, s.devices.Insert(&device)
}

func (s *AuthService) UpdateDevice(id, pushToken string) error {
	return s.devices.UpdateToken(id, pushToken)
}

func (s *AuthService) AuthorizeDevice(token string) (models.Device, error) {
	device, err := s.devices.GetByToken(token)
	if err != nil {
		return device, err
	}

	if err := s.devices.UpdateLastSeen(device.ID); err != nil {
		s.logger.Error("can't update last seen", zap.Error(err))
	}

	return device, nil
}

func (s *AuthService) AuthorizeUser(username, password string) (models.User, error) {
	user, err := s.users.GetByLogin(username)
	if err != nil {
		return user, err
	}

	return user, crypto.CompareBCryptHash(user.PasswordHash, password)
}

type AuthServiceParams struct {
	fx.In

	Users   *repositories.UsersRepository
	Devices *repositories.DevicesRepository

	Logger *zap.Logger
}

func NewAuthService(params AuthServiceParams) *AuthService {
	idgen, _ := nanoid.Standard(21)

	return &AuthService{
		users:   params.Users,
		devices: params.Devices,
		logger:  params.Logger.Named("AuthService"),
		idgen:   idgen,
	}
}
