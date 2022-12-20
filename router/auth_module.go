package router

import (
	"github.com/gofiber/fiber/v2"
	"mindmap-go/app/controllers"
)

type AuthRouter struct {
	App        fiber.Router
	Controller *controllers.Controller
}

// Router methods
func NewAuthRouter(fiber *fiber.App, controller *controllers.Controller) *UserRouter {
	return &UserRouter{
		App:        fiber,
		Controller: controller,
	}
}

func (r *UserRouter) RegisterAuthRoutes() {
	// Define controllers
	authController := r.Controller.Auth

	// Define routes
	r.App.Route("/auth", func(router fiber.Router) {
		router.Post("/register", authController.Register)
		router.Post("/login", authController.Login)
	})
}
