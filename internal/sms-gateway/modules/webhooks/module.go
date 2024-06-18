package webhooks

import (
	"github.com/capcom6/go-infra-fx/db"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module(
	"webhooks",
	fx.Decorate(func(log *zap.Logger) *zap.Logger {
		return log.Named("webhooks")
	}),
	fx.Provide(NewRepository, fx.Private),
	fx.Provide(
		NewService,
	),
)

func init() {
	db.RegisterMigration(Migrate)
}
