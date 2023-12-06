package models

import (
	"github.com/capcom6/sms-gateway/internal/infra/db"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"models",
	fx.Provide(
		db.AsMigration(NewMigration),
		GetGooseStorage,
	),
)
