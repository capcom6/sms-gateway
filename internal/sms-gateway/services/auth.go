package services

import (
	"fmt"

	"github.com/capcom6/sms-gateway/internal/sms-gateway/models"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/repositories"
	"github.com/capcom6/sms-gateway/pkg/crypto"
	"github.com/jaevor/go-nanoid"
)

type AuthService struct {
	users   *repositories.UsersRepository
	devices *repositories.DevicesRepository

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
	return s.devices.GetByToken(token)
}

func (s *AuthService) AuthorizeUser(username, password string) (models.User, error) {
	user, err := s.users.GetByLogin(username)
	if err != nil {
		return user, err
	}

	return user, crypto.CompareBCryptHash(user.PasswordHash, password)
}

func NewAuthService(users *repositories.UsersRepository, devices *repositories.DevicesRepository) *AuthService {
	idgen, _ := nanoid.Standard(21)

	return &AuthService{
		users:   users,
		devices: devices,
		idgen:   idgen,
	}
}
