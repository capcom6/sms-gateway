package handlers

import (
	microbase "bitbucket.org/soft-c/gomicrobase"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

type thirdPartyHandler struct {
	microbase.Handler
}

func (h *thirdPartyHandler) postMessage(c *fiber.Ctx) error {
	return fiber.ErrNotImplemented
}

func (h *thirdPartyHandler) register(router fiber.Router) {
	router.Use(basicauth.New(basicauth.Config{
		Authorizer: func(username string, password string) bool {
			return len(username) > 0 && len(password) > 0
		},
	}))

	router.Post("/message", h.postMessage)
}

func newThirdPartyHandler() *thirdPartyHandler {
	return &thirdPartyHandler{}
}
