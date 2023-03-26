package bootstrap

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"mindmap-go/internal/database"
	"mindmap-go/router"
	"mindmap-go/utils/config"
	"mindmap-go/utils/response"
	"os"
	"runtime"
	"time"
)

func NewFiber(cfg *config.Config) *fiber.App {
	// Setup Webserver
	app := fiber.New(fiber.Config{
		ServerHeader:          cfg.App.Name,
		AppName:               cfg.App.Name,
		Prefork:               cfg.App.Prefork,
		ErrorHandler:          response.ErrorHandler,
		IdleTimeout:           cfg.App.IdleTimeout * time.Second,
		EnablePrintRoutes:     cfg.App.PrintRoutes,
		DisableStartupMessage: true,
	})

	// Pass production config to check it
	response.IsProduction = cfg.App.Production

	return app
}

func Start(lifecycle fx.Lifecycle, cfg *config.Config, fiber *fiber.App, router *router.Router, database *database.Database, log *zap.Logger) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				// Register middlewares & routes
				//middlewares.Register()
				router.Register()

				// Custom Startup Messages
				host, port := config.ParseAddr(cfg.App.Port)
				if host == "" {
					if fiber.Config().Network == "tcp6" {
						host = "[::1]"
					} else {
						host = "0.0.0.0"
					}
				}

				// Information message
				log.Info(fiber.Config().AppName + " is running at the moment!")
				log.Debug("", zap.Bool("isProduction", cfg.App.Production))
				// Debug information
				if !cfg.App.Production {
					prefork := "Enabled"
					procs := runtime.GOMAXPROCS(0)
					if !cfg.App.Prefork {
						procs = 1
						prefork = "Disabled"
					}

					log.Info("", zap.String("Host", host))
					log.Debug("", zap.String("Port", port))
					log.Debug("", zap.String("Prefork", prefork))
					log.Debug("", zap.Uint32("Handlers", fiber.HandlersCount()))
					log.Debug("", zap.Int("Processes", procs))
					log.Debug("", zap.Int("PID", os.Getpid()))
				}

				go func() {
					if err := fiber.Listen(cfg.App.Port); err != nil {
						log.Error(fmt.Sprintf("An unknown error occurred when to run server: %s", err.Error()))
					}
				}()

				database.OpenConnection()

				return nil
			},
			OnStop: func(ctx context.Context) error {
				log.Info("Shutting down the app...")
				if err := fiber.Shutdown(); err != nil {
					log.Panic(err.Error())
				}

				log.Info("Running cleanup tasks...")
				log.Info("1- Shutdown the database")
				database.CloseConnection()
				log.Info(fmt.Sprintf("%s was successfuly shutdown.", cfg.App.Name))
				log.Info("This was a triumph")
				log.Info("I'm making a note here: HUGE SUCCESS.")

				return nil
			},
		},
	)
}
