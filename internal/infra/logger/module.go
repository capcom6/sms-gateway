package logger

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module(
	"logger",
	fx.Provide(New),
	fx.Invoke(func(logger *zap.Logger) {
		zap.RedirectStdLog(logger)
	}),
)
