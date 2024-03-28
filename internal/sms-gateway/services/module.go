package services

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module(
	"services",
	fx.Decorate(func(log *zap.Logger) *zap.Logger {
		return log.Named("services")
	}),
	fx.Provide(
		NewDevicesService,
	),
)
