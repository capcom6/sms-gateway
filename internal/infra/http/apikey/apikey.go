package apikey

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

// New creates a new middleware handler
func New(config Config) fiber.Handler {
	// Set default config
	cfg := configDefault(config)

	// Return new handler
	return func(c *fiber.Ctx) error {
		// Don't execute middleware if Next returns true
		if cfg.Next != nil && cfg.Next(c) {
			return c.Next()
		}

		// Get authorization header
		auth := c.Get(fiber.HeaderAuthorization)

		// Check if the header contains content besides "bearer".
		if len(auth) <= 6 || strings.ToLower(auth[:6]) != "bearer" {
			return cfg.Unauthorized(c)
		}

		token := auth[7:]

		// Check if the token is empty.
		if len(token) == 0 {
			return cfg.Unauthorized(c)
		}

		if cfg.Authorizer(token) {
			c.Locals(cfg.ContextToken, token)
			return c.Next()
		}

		// Authentication failed
		return cfg.Unauthorized(c)
	}
}
