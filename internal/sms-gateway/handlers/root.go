package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type rootHandler struct {
	healthHandler *healthHandler
}

func (h *rootHandler) Register(app *fiber.App) {
	h.healthHandler.Register(app)
	app.Static("/", "static")
}

func newRootHandler(healthHandler *healthHandler) *rootHandler {
	return &rootHandler{
		healthHandler: healthHandler,
	}
}
