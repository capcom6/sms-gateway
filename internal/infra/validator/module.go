package validator

import "go.uber.org/fx"

var Module = fx.Module(
	"validator",
	fx.Provide(New),
)
