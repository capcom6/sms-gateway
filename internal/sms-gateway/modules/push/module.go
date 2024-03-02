package push

import (
	"context"

	"github.com/capcom6/sms-gateway/internal/sms-gateway/modules/push/fcm"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module(
	"push",
	fx.Decorate(func(log *zap.Logger) *zap.Logger {
		return log.Named("push")
	}),
	fx.Provide(
		func(cfg Config, lc fx.Lifecycle) (client, error) {
			client, err := fcm.New(cfg.ClientOptions)
			if err != nil {
				return nil, err
			}

			lc.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					return client.Open(ctx)
				},
				OnStop: func(ctx context.Context) error {
					return client.Close(ctx)
				},
			})

			return client, nil
		},
	),
	fx.Provide(
		New,
	),
)
