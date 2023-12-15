package tasks

import (
	"context"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module(
	"tasks",
	fx.Decorate(func(log *zap.Logger) *zap.Logger {
		return log.Named("tasks")
	}),
	fx.Provide(
		NewHashingTask,
		fx.Private,
	),
	fx.Invoke(
		func(lc fx.Lifecycle, task *HashingTask) error {
			ctx, cancel := context.WithCancel(context.Background())

			lc.Append(fx.Hook{
				OnStart: func(_ context.Context) error {
					go task.Run(ctx)
					return nil
				},
				OnStop: func(_ context.Context) error {
					cancel()
					return nil
				},
			})

			return nil
		},
	),
)
