package db

import (
	"bitbucket.org/capcom6/smsgatewaybackend/internal/infra/cli"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"db",
	fx.Provide(
		New,
		cli.AsCommand(NewCommandMigrate),
	),
)
