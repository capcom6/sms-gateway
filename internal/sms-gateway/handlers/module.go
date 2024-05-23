package handlers

import (
	"github.com/capcom6/go-infra-fx/http"
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
		fx.Private,
	),
)
