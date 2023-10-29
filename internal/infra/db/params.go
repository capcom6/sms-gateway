package db

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Params struct {
	fx.In

	Logger *zap.Logger
	Config Config
	// Migrations []Migrator `group:"migrations"`
	LC fx.Lifecycle
}
