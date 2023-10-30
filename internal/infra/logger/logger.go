package logger

import (
	"context"
	"errors"
	"os"
	"syscall"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func New(lc fx.Lifecycle) (*zap.Logger, error) {
	isDebug := os.Getenv("DEBUG") != ""

	logConfig := zap.NewProductionConfig()
	if isDebug {
		logConfig = zap.NewDevelopmentConfig()
	}

	l, err := logConfig.Build()
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
