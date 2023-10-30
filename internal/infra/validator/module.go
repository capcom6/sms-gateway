package validator

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module(
	"validator",
	fx.Decorate(func(log *zap.Logger) *zap.Logger {
		return log.Named("validator")
	}),
	fx.Provide(New),
)
