package handlers

import (
	"bitbucket.org/capcom6/smsgatewaybackend/internal/smsgateway/repositories"
	"bitbucket.org/capcom6/smsgatewaybackend/internal/smsgateway/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Register(router fiber.Router, db *gorm.DB) error {
	users := repositories.NewUsersRepository(db)
	devices := repositories.NewDevicesRepository(db)
	messages := repositories.NewMessagesRepository(db)

	validator := validator.New()
	authSvc := services.NewAuthService(users, devices)
	messagesSvc := services.NewMessagesService(messages)

	newMobileHandler(validator, authSvc, messagesSvc).register(router.Group("/mobile/v1"))
	newThirdPartyHandler(validator, authSvc, messagesSvc).register(router.Group("/3rdparty/v1"))

	return nil
}
