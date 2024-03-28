package messages

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// TODO: merge service and hashing task configs
// TODO: run hashing task inside service

var Module = fx.Module(
	"messages",
	fx.Decorate(func(log *zap.Logger) *zap.Logger {
		return log.Named("messages")
	}),
	fx.Provide(NewService),
	fx.Provide(NewHashingTask, fx.Private),
)
