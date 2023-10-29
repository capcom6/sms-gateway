package apikey

import "github.com/gofiber/fiber/v2"

// Config defines the config for middleware.
type Config struct {
	// Next defines a function to skip this middleware when returned true.
	//
	// Optional. Default: nil
	Next func(c *fiber.Ctx) bool

	// Authorizer defines a function you can pass
	// to check the credentials however you want.
	// It will be called with a token
	// and is expected to return true or false to indicate
	// that the credentials were approved or not.
	//
	// Optional. Default: nil.
	Authorizer func(string) bool

	// Unauthorized defines the response body for unauthorized responses.
	// By default it will return with a 401 Unauthorized and the correct WWW-Auth header
	//
	// Optional. Default: nil
	Unauthorized fiber.Handler

	// ContextToken is the key to store the token in Locals
	//
	// Optional. Default: "token"
	ContextToken string
}

// ConfigDefault is the default config
var ConfigDefault = Config{
	Next:         nil,
	Authorizer:   nil,
	Unauthorized: nil,
	ContextToken: "token",
}

// Helper function to set default values
func configDefault(config ...Config) Config {
	// Return default config if nothing provided
	if len(config) < 1 {
		return ConfigDefault
	}

	// Override default config
	cfg := config[0]

	// Set default values
	if cfg.Next == nil {
		cfg.Next = ConfigDefault.Next
	}
	if cfg.Authorizer == nil {
		cfg.Authorizer = func(token string) bool {
			return true
		}
	}
	if cfg.Unauthorized == nil {
		cfg.Unauthorized = func(c *fiber.Ctx) error {
			return fiber.ErrUnauthorized
		}
	}
	if cfg.ContextToken == "" {
		cfg.ContextToken = ConfigDefault.ContextToken
	}
	return cfg
}
