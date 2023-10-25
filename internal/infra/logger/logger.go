package logger

import (
	"context"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func New(lc fx.Lifecycle) (*zap.Logger, error) {
	l, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return l.Sync()
		},
	})

	return l, nil
}
