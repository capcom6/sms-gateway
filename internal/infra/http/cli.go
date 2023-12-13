package http

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type RunServerParams struct {
	fx.In

	Config Config
	App    *fiber.App
	Logger *zap.Logger
	LC     fx.Lifecycle
}

func Run(params RunServerParams) error {
	go func() {
		params.Logger.Info("Starting server...")

		err := params.App.Listen(params.Config.Listen)
		if err != nil {
			params.Logger.Error("Error starting server", zap.Error(err))
		}
	}()

	params.LC.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return params.App.ShutdownWithContext(ctx)
		},
	})

	return nil
}
