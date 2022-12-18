package routes

import (
	"fiber-boilerplate/database"

	"github.com/gofiber/fiber/v2"
)

func RegisterAPI(api fiber.Router, db *database.Database) {
	RegisterArticleRoutes()
}

type ArticleRouter struct {
	App        fiber.Router
	Controller *controller.Controller
}

func (r *ArticleRouter) RegisterArticleRoutes() {
	// Define controllers
	articleController := r.Controller.Article

	// Define routes
	r.App.Route("/articles", func(router fiber.Router) {
		router.Get("/", articleController.Index)
		router.Get("/:id", articleController.Show)
		router.Post("/", articleController.Store)
		router.Patch("/:id", articleController.Update)
		router.Delete("/:id", articleController.Destroy)
	})
}
