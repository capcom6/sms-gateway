package http

import (
	"github.com/capcom6/sms-gateway/internal/infra/cli"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"http",
	fx.Provide(
		New,
		cli.AsCommand(NewRunServer),
	),
)
