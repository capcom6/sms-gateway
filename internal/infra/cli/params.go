package cli

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Params struct {
	fx.In

	Logger   *zap.Logger
	Commands []Command `group:"commands"`
	Shut     fx.Shutdowner
}
