package http

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Params struct {
	fx.In

	Config       Config
	Logger       *zap.Logger
	ApiHandlers  []ApiHanlder  `group:"api-routes"`
	RootHandlers []RootHanlder `group:"root-routes"`
	LC           fx.Lifecycle
}
