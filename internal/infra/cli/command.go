package cli

import "go.uber.org/fx"

type Command interface {
	Cmd() string
	Run(args ...string) error
}

func AsCommand(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(Command)),
		fx.ResultTags(`group:"commands"`),
	)
}
