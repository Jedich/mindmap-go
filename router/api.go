package router

import (
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	App        fiber.Router
	UserRouter *UserRouter
}

func NewRouter(fiber *fiber.App, articleRouter *UserRouter) *Router {
	return &Router{
		App:        fiber,
		UserRouter: articleRouter,
	}
}

// Register routes
func (r *Router) Register() {
	// Test Routes
	r.App.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Pong! 👋")
	})

	// Register routes of modules
	r.UserRouter.RegisterUserRoutes()
}
