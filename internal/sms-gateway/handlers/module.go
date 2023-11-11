package handlers

import (
	"github.com/capcom6/sms-gateway/internal/infra/http"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"handlers",
	fx.Provide(
		http.AsRootHandler(newRootHandler),
		http.AsApiHandler(newThirdPartyHandler),
		http.AsApiHandler(newMobileHandler),
	),
)
