package models

import (
	"bitbucket.org/capcom6/smsgatewaybackend/internal/infra/db"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"models",
	fx.Provide(
		db.AsMigration(NewMigration),
	),
)
