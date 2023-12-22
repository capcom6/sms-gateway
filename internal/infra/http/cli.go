package http

import (
	"context"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

type RunServerParams struct {
	fx.In

	Server *Server
	Logger *zap.Logger
	LC     fx.Lifecycle
	Shut   fx.Shutdowner
}

func Run(params RunServerParams) error {
	params.LC.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				if err := params.Server.Start(); err != nil {
					params.Logger.Error("Error starting server", zap.Error(err))
					params.Shut.Shutdown()
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return params.Server.Stop(ctx)
		},
	})

	return nil
}
