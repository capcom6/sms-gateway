package devices

import (
	"github.com/capcom6/sms-gateway/internal/sms-gateway/modules/cleaner"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type FxResult struct {
	fx.Out

	Service   *Service
	AsCleaner cleaner.Cleanable `group:"cleaners"`
}

var Module = fx.Module(
	"devices",
	fx.Decorate(func(log *zap.Logger) *zap.Logger {
		return log.Named("devices")
	}),
	fx.Provide(
		newDevicesRepository,
		fx.Private,
	),
	fx.Provide(func(p ServiceParams) FxResult {
		svc := NewService(p)
		return FxResult{
			Service:   svc,
			AsCleaner: svc,
		}
	}),
)
