package utils

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
)

func IsEnabled(key bool) func(c *fiber.Ctx) bool {
	if key {
		return nil
	}

	return func(c *fiber.Ctx) bool { return true }
}

func ReadEnv(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("%s not set\n", key))
	}
	return val
}
