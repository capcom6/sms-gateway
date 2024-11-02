package cleaner

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func AsCleanable(src any) any {
	return fx.Annotate(
		src,
		fx.As(new(Cleanable)),
		fx.ResultTags(`group:"cleaners"`),
	)
}

type Params struct {
	fx.In

	Cleanables []Cleanable `group:"cleaners"`

	Logger *zap.Logger
}

func NewFx(p Params) *Service {
	return New(p.Cleanables, p.Logger)
}

var Module = fx.Module(
	"cleaner",
	fx.Decorate(func(log *zap.Logger) *zap.Logger {
		return log.Named("cleaner")
	}),
	fx.Provide(
		NewFx,
	),
)
