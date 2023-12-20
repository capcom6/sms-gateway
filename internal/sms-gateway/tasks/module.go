package tasks

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module(
	"tasks",
	fx.Decorate(func(log *zap.Logger) *zap.Logger {
		return log.Named("tasks")
	}),
	fx.Provide(NewHashingTask),
)
