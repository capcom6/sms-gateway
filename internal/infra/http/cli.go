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

type RunServer struct {
	Config Config
	App    *fiber.App
	Logger *zap.Logger
	LC     fx.Lifecycle
}

func NewRunServer(params RunServerParams) *RunServer {
	return &RunServer{
		Config: configDefault(params.Config),
		App:    params.App,
		Logger: params.Logger,
		LC:     params.LC,
	}
}

func (c *RunServer) Cmd() string {
	return "runserver"
}

func (c *RunServer) Run(args ...string) error {
	go func() {
		err := c.App.Listen(c.Config.Listen)
		if err != nil {
			c.Logger.Error("Error starting server", zap.Error(err))
		}
	}()

	c.LC.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return c.App.ShutdownWithContext(ctx)
		},
	})

	return nil
}
