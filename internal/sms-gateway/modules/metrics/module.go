package metrics

import (
	"github.com/capcom6/go-infra-fx/http"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module(
	"metrics",
	fx.Decorate(func(log *zap.Logger) *zap.Logger {
		return log.Named("metrics")
	}),
	fx.Provide(
		http.AsRootHandler(newHttpHandler),
	),
)
