package auth

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module(
	"auth",
	fx.Decorate(func(log *zap.Logger) *zap.Logger {
		return log.Named("auth")
	}),
	fx.Provide(New),
)
