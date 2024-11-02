package messages

import (
	"github.com/capcom6/sms-gateway/internal/sms-gateway/modules/cleaner"
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
	// fx.Provide(cleaner.AsCleanable(NewService)),
	// fx.Provide(fx.Annotate(NewService, fx.ResultTags(`group:"cleaners"`))),
	fx.Provide(func(p ServiceParams) FxResult {
		svc := NewService(p)
		return FxResult{
			Service:   svc,
			AsCleaner: svc,
		}
	}),
	fx.Provide(NewHashingTask, fx.Private),
)
