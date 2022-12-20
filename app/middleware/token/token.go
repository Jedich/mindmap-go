// Simple token middleware as example

package token

import (
	"github.com/gofiber/fiber/v2"
)

type Config struct {
	Next       func(c *fiber.Ctx) bool
	Token      string
	HeaderName string
}

var ConfigDefault = Config{
	Next:       nil,
	Token:      "Default",
	HeaderName: "X-Token-Middleware",
}

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
	if cfg.Token == "" {
		cfg.Token = ConfigDefault.Token
	}
	if cfg.HeaderName == "" {
		cfg.HeaderName = ConfigDefault.HeaderName
	}

	return cfg
}

// New creates a new middleware handler
func New() fiber.Handler {
	// Return new handler
	return func(c *fiber.Ctx) error {
		return nil
	}
}
