package router

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"mindmap-go/app/controllers"
	"mindmap-go/app/repository"
	"mindmap-go/app/services"
)

type UserRouter struct {
	App        fiber.Router
	Controller *controllers.Controller
}

// NewUserModule add all dependencies
var NewUserModule = fx.Options(

	// Register Repository & Service
	fx.Provide(repository.NewAccountRepository),
	fx.Provide(repository.NewUserRepository),
	fx.Provide(services.NewUserService),

	// Regiser Controller
	fx.Provide(controllers.NewController),

	// Register Router
	fx.Provide(NewUserRouter),
)

// Router methods
func NewUserRouter(fiber *fiber.App, controller *controllers.Controller) *UserRouter {
	return &UserRouter{
		App:        fiber,
		Controller: controller,
	}
}

func (r *UserRouter) RegisterUserRoutes() {
	// Define controllers
	userController := r.Controller.User

	// Define routes
	r.App.Route("/users", func(router fiber.Router) {
		router.Get("/", userController.Index)
		router.Get("/:id", userController.Show)
		router.Post("/new", userController.Store)
		router.Patch("/:id", userController.Update)
		router.Delete("/:id", userController.Destroy)
	})
}
