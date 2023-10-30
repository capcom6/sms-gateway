package cli

import (
	"os"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

var DefaultCommand = ""

var Module = fx.Module(
	"cli",
	fx.Decorate(func(log *zap.Logger) *zap.Logger {
		return log.Named("cli")
	}),
	fx.Invoke(func(params Params) error {
		cmd := DefaultCommand
		args := []string{}
		if len(os.Args) > 1 {
			cmd = os.Args[1]
			args = os.Args[2:]
		}

		for _, v := range params.Commands {
			if v.Cmd() != cmd {
				continue
			}

			return v.Run(args...)
		}
		params.Logger.Info("Command is not supported", zap.String("command", cmd))
		return params.Shut.Shutdown()
	}),
)
