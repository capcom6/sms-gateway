package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type rootHandler struct {
}

func (h *rootHandler) Register(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("https://github.com/capcom6/android-sms-gateway")
	})
}

func newRootHandler() *rootHandler {
	return &rootHandler{}
}
