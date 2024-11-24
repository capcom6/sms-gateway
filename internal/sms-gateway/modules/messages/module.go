package messages

import (
	"github.com/android-sms-gateway/server/internal/sms-gateway/modules/cleaner"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// TODO: merge service and hashing task configs
// TODO: run hashing task inside service

type FxResult struct {
	fx.Out

	Service   *Service
	AsCleaner cleaner.Cleanable `group:"cleaners"`
}

var Module = fx.Module(
	"messages",
	fx.Decorate(func(log *zap.Logger) *zap.Logger {
		return log.Named("messages")
	}),
	fx.Provide(func(p ServiceParams) FxResult {
		svc := NewService(p)
		return FxResult{
			Service:   svc,
			AsCleaner: svc,
		}
	}),
	fx.Provide(newRepository),
	fx.Provide(NewHashingTask, fx.Private),
)
