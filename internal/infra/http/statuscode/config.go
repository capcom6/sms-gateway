package statuscode

import "github.com/gofiber/fiber/v2"

// Config defines the config for middleware.
type Config struct {
	// Next defines a function to skip this middleware when returned true.
	//
	// Optional. Default: nil
	Next func(c *fiber.Ctx) bool
	// Response status code
	//
	// Optional. Default: 404
	StatusCode int
	// Response status message
	//
	// Optional. Default: Not Found
	StatusMessage string
}

// ConfigDefault is the default config
var ConfigDefault = Config{
	Next:          nil,
	StatusCode:    404,
	StatusMessage: "Not Found",
}

// Helper function to set default values
func configDefault(config ...Config) Config {
	// Return default config if nothing provided
	if len(config) < 1 {
		return ConfigDefault
	}

	// Override default config
	cfg := config[0]

	if cfg.StatusCode == 0 {
		cfg.StatusCode = fiber.StatusNotFound
	}
	if cfg.StatusMessage == "" {
		cfg.StatusMessage = "Not Found"
	}

	return cfg
}
