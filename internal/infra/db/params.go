package db

import (
	"database/sql"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Params struct {
	fx.In

	Logger *zap.Logger
	Config Config
	SQL    *sql.DB
	LC     fx.Lifecycle
}
