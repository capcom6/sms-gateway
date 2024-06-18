package devices

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module(
	"devices",
	fx.Decorate(func(log *zap.Logger) *zap.Logger {
		return log.Named("devices")
	}),
	fx.Provide(
		newDevicesRepository,
		fx.Private,
	),
	fx.Provide(
		NewService,
	),
)
