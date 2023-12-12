package main

import (
	_ "go.uber.org/automaxprocs"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"mindmap-go/internal/bootstrap"
	"mindmap-go/internal/database"
	"mindmap-go/router"
	"mindmap-go/utils/config"
)

func main() {
	fx.New(
		// Provide patterns
		fx.Provide(config.NewConfig),
		fx.Provide(bootstrap.NewLogger),
		fx.Provide(bootstrap.NewFiber),
		fx.Provide(database.NewDatabase),
		fx.Provide(router.NewRouter),

		// Provide modules
		router.NewUserModule,

		// Start Application
		fx.Invoke(bootstrap.Start),

		// Define logger
		fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: logger}
		}),
	).Run()
}
