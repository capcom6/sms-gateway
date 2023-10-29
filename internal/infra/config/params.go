package config

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Param struct {
	fx.In

	Logger *zap.Logger
	Config any `name:"config:source"`
}
