package handlers

import (
	"bitbucket.org/capcom6/smsgatewaybackend/internal/infra/http"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"handlers",
	fx.Provide(
		http.AsApiHandler(newThirdPartyHandler),
		http.AsApiHandler(newMobileHandler),
	),
)
