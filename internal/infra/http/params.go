package http

import "go.uber.org/fx"

type Params struct {
	fx.In

	Config      Config
	ApiHandlers []ApiHanlder `group:"api-routes"`
	LC          fx.Lifecycle
}
