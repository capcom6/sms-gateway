package handlers

import (
	"github.com/capcom6/go-infra-fx/http"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/handlers/devices"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/handlers/logs"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/handlers/webhooks"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module(
	"handlers",
	fx.Decorate(func(log *zap.Logger) *zap.Logger {
		return log.Named("handlers")
	}),
	fx.Provide(
		http.AsRootHandler(newRootHandler),
		http.AsApiHandler(newThirdPartyHandler),
		http.AsApiHandler(newMobileHandler),
		http.AsApiHandler(newUpstreamHandler),
	),
	fx.Provide(
		newHealthHandler,
		webhooks.NewThirdPartyController,
		webhooks.NewMobileController,
		devices.NewThirdPartyController,
		logs.NewThirdPartyController,
		fx.Private,
	),
)
