package logger

import (
	"context"
	"errors"
	"syscall"

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
			if err := l.Sync(); !errors.Is(err, syscall.ENOTTY) {
				return err
			}
			return nil
		},
	})

	return l, nil
}
