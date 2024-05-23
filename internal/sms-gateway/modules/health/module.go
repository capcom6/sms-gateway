package health

import (
	"github.com/capcom6/go-infra-fx/cli"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module(
	"health",
	fx.Decorate(func(log *zap.Logger) *zap.Logger {
		return log.Named("health")
	}),
	fx.Provide(
		AsHealthProvider(NewDBProvider),
		fx.Private,
	),
	fx.Provide(
		NewService,
	),
)

func init() {
	cli.Register("health", testHealth)
}
