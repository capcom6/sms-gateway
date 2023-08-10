package handlers

import (
	"bitbucket.org/capcom6/smsgatewaybackend/internal/config"
	"bitbucket.org/capcom6/smsgatewaybackend/internal/smsgateway/repositories"
	"bitbucket.org/capcom6/smsgatewaybackend/internal/smsgateway/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Register(router fiber.Router, db *gorm.DB) error {
	cfg := config.GetConfig()

	users := repositories.NewUsersRepository(db)
	devices := repositories.NewDevicesRepository(db)
	messages := repositories.NewMessagesRepository(db)

	validator := validator.New()
	authSvc := services.NewAuthService(users, devices)
	pushSvc := services.NewPushService(cfg.FCM.CredentialsJSON)
	messagesSvc := services.NewMessagesService(pushSvc, messages)

	newMobileHandler(validator, authSvc, messagesSvc).register(router.Group("/mobile/v1"))
	newThirdPartyHandler(validator, authSvc, messagesSvc).register(router.Group("/3rdparty/v1"))

	return nil
}
