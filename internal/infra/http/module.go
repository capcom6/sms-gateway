package http

import (
	"github.com/capcom6/sms-gateway/internal/infra/cli"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module(
	"http",
	fx.Decorate(func(log *zap.Logger) *zap.Logger {
		return log.Named("http")
	}),
	fx.Provide(
		New,
		cli.AsCommand(NewRunServer),
	),
)
