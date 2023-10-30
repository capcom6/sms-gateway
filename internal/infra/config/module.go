package config

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module(
	"config",
	fx.Decorate(func(log *zap.Logger) *zap.Logger {
		return log.Named("config")
	}),
	fx.Provide(
		fx.Annotate(
			New,
			fx.ResultTags(`name:"config:result"`),
		),
	),
)
