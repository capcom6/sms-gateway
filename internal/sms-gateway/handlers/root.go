package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type rootHandler struct {
	healthHandler *healthHandler
}

func (h *rootHandler) Register(app *fiber.App) {
	app.Use(func(c *fiber.Ctx) error {
		if c.Path() == "/api" {
			return c.Redirect("/api/", fiber.StatusMovedPermanently)
		}

		return c.Next()
	})

	h.healthHandler.Register(app)
	app.Static("/api", "api")
}

func newRootHandler(healthHandler *healthHandler) *rootHandler {
	return &rootHandler{
		healthHandler: healthHandler,
	}
}
