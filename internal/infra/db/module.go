package db

import (
	"github.com/capcom6/sms-gateway/internal/infra/cli"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"db",
	fx.Provide(
		New,
		cli.AsCommand(NewCommandMigrate),
	),
)
