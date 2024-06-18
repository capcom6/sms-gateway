package auth

import (
	"github.com/capcom6/sms-gateway/internal/sms-gateway/models"
	"github.com/gofiber/fiber/v2"
)

func WithUser(handler func(models.User, *fiber.Ctx) error) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return handler(c.Locals("user").(models.User), c)
	}
}

func WithDevice(handler func(models.Device, *fiber.Ctx) error) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return handler(c.Locals("device").(models.Device), c)
	}
}
