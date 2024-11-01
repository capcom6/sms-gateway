package handlers

import (
	"net/http"

	"github.com/capcom6/sms-gateway/pkg/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
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
	app.Use("/api", filesystem.New(filesystem.Config{
		Root:       http.FS(swagger.Docs),
		PathPrefix: "docs",
		MaxAge:     1 * 24 * 60 * 60,
	}), func(c *fiber.Ctx) error {
		// The filesystem middleware set 404 status before next, so we need to override it
		return c.Status(fiber.StatusOK).Next()
	})
}

func newRootHandler(healthHandler *healthHandler) *rootHandler {
	return &rootHandler{
		healthHandler: healthHandler,
	}
}
