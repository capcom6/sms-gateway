package statuscode

import "github.com/gofiber/fiber/v2"

// New creates a new middleware handler
func New(config ...Config) fiber.Handler {
	cfg := configDefault(config...)

	return func(c *fiber.Ctx) error {
		// Don't execute middleware if Next returns true
		if cfg.Next != nil && cfg.Next(c) {
			return c.Next()
		}

		return c.Status(cfg.StatusCode).SendString(cfg.StatusMessage)
	}
}
