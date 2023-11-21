package db

import (
	"github.com/capcom6/sms-gateway/internal/infra/cli"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module(
	"db",
	fx.Decorate(func(log *zap.Logger) *zap.Logger {
		return log.Named("db")
	}),
	fx.Provide(
		New,
		NewSQL,
		cli.AsCommand(NewCommandMigrate),
		cli.AsCommand(NewGooseMigrate),
	),
)
