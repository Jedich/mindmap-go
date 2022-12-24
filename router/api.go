package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	jwtware "github.com/gofiber/jwt/v3"
	"mindmap-go/utils"
)

type Router struct {
	App        fiber.Router
	UserRouter *UserRouter
	MapRouter  *MapRouter
	CardRouter *CardRouter
}

func NewRouter(fiber *fiber.App, userRouter *UserRouter, mapRouter *MapRouter, cardRouter *CardRouter) *Router {
	return &Router{
		App:        fiber,
		UserRouter: userRouter,
		MapRouter:  mapRouter,
		CardRouter: cardRouter,
	}
}

// Register routes
func (r *Router) Register() {
	r.App.Use(logger.New())

	// Routes, unrestricted access
	r.App.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Pong! 👋")
	})

	// Register auth routes
	r.UserRouter.RegisterAuthRoutes()

	// JWT Middleware
	r.App.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(utils.ReadEnv("JWT_SECRET")),
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return &utils.UnauthorizedEntryError{}
		},
	}))

	// Register routes of modules, restricted access
	r.UserRouter.RegisterUserRoutes()
	r.MapRouter.RegisterMapRoutes()
	r.CardRouter.RegisterCardRoutes()
}
