package smsgateway

import (
	appconfig "github.com/capcom6/sms-gateway/internal/config"
	"github.com/capcom6/sms-gateway/internal/infra/cli"
	"github.com/capcom6/sms-gateway/internal/infra/config"
	"github.com/capcom6/sms-gateway/internal/infra/db"
	"github.com/capcom6/sms-gateway/internal/infra/http"
	"github.com/capcom6/sms-gateway/internal/infra/logger"
	"github.com/capcom6/sms-gateway/internal/infra/validator"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/handlers"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/models"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/repositories"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/services"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"server",
	cli.Module,
	appconfig.Module,
	config.Module,
	logger.Module,
	http.Module,
	validator.Module,
	handlers.Module,
	services.Module,
	repositories.Module,
	models.Module,
	db.Module,
)

func Run() {
	cli.DefaultCommand = "http:run"
	fx.New(
		Module,
	).Run()
}