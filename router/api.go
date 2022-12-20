package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	jwtware "github.com/gofiber/jwt/v3"
	"mindmap-go/utils"
	"mindmap-go/utils/config"
)

type Middleware struct {
	App *fiber.App
	Cfg *config.Config
}

func NewMiddleware(app *fiber.App, cfg *config.Config) *Middleware {
	return &Middleware{
		App: app,
		Cfg: cfg,
	}
}

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
	r.App.Use(logger.New())

	// Test Routes
	r.App.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Pong! 👋")
	})

	// Register routes of modules
	r.UserRouter.RegisterAuthRoutes()

	// JWT Middleware
	r.App.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(utils.ReadEnv("JWT_SECRET")),
	}))

	// Register routes of modules
	r.UserRouter.RegisterUserRoutes()
}
