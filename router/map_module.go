package router

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"mindmap-go/app/controllers"
	"mindmap-go/app/repository"
	"mindmap-go/app/services"
)

type MapRouter struct {
	App           fiber.Router
	MapController controllers.MapController
}

// NewMapModule add all dependencies
var NewMapModule = fx.Options(

	// Register Repository & Service
	fx.Provide(repository.NewMapRepository),
	fx.Provide(services.NewMapService),

	// Regiser Controller
	fx.Provide(controllers.NewMapController),

	// Register Router
	fx.Provide(NewMapRouter),
)

// NewMapRouter methods
func NewMapRouter(fiber *fiber.App, controller controllers.MapController) *MapRouter {
	return &MapRouter{
		App:           fiber,
		MapController: controller,
	}
}

func (r *MapRouter) RegisterMapRoutes() {
	// Define controllers
	mapController := r.MapController

	// Define routes
	r.App.Route("/maps", func(router fiber.Router) {
		router.Get("/", mapController.Index)
		router.Get("/:id", mapController.Show)
		router.Post("/", mapController.Store)
		router.Patch("/:id", mapController.Update)
		router.Delete("/:id", mapController.Destroy)
	})
}
