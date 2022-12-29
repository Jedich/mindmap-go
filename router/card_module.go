package router

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"mindmap-go/app/controllers"
	"mindmap-go/app/repository"
	"mindmap-go/app/services"
)

type CardRouter struct {
	App            fiber.Router
	CardController controllers.CardController
}

// NewMapModule add all dependencies
var NewCardModule = fx.Options(

	// Register Repository & Service
	fx.Provide(repository.NewCardRepository),
	fx.Provide(services.NewCardService),

	// Regiser Controller
	fx.Provide(controllers.NewCardController),

	// Register Router
	fx.Provide(NewCardRouter),
)

// NewMapRouter methods
func NewCardRouter(fiber *fiber.App, controller controllers.CardController) *CardRouter {
	return &CardRouter{
		App:            fiber,
		CardController: controller,
	}
}

func (r *CardRouter) RegisterCardRoutes() {
	// Define controllers
	cardController := r.CardController

	// Define routes
	r.App.Route("/cards", func(router fiber.Router) {
		router.Get("/:id", cardController.Index)
		router.Post("/", cardController.Store)
		router.Patch("/", cardController.Update)
		router.Delete("/:id", cardController.Destroy)
	})
}
