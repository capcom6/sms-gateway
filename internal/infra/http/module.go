package http

import (
	"bitbucket.org/capcom6/smsgatewaybackend/internal/infra/cli"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"http",
	fx.Provide(
		New,
		cli.AsCommand(NewRunServer),
	),
)
