package cli

import (
	"os"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

var DefaultCommand = ""

func GetModule() fx.Option {
	cmd := DefaultCommand
	args := []string{}
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}

	executor, ok := commands[cmd]
	if !ok {
		return fx.Invoke(func(logger *zap.Logger, shut fx.Shutdowner) error {
			logger.Error("Command is not supported", zap.String("cmd", cmd))
			return shut.Shutdown()
		})
	}

	return fx.Module(
		"cli",
		fx.Supply(Args(args)),
		fx.Invoke(executor),
	)
}
