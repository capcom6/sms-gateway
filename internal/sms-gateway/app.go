package smsgateway

import (
	appconfig "github.com/capcom6/sms-gateway/internal/config"
	"github.com/capcom6/sms-gateway/internal/infra/cli"
	"github.com/capcom6/sms-gateway/internal/infra/db"
	"github.com/capcom6/sms-gateway/internal/infra/http"
	"github.com/capcom6/sms-gateway/internal/infra/logger"
	"github.com/capcom6/sms-gateway/internal/infra/validator"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/handlers"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/models"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/repositories"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/services"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/tasks"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Module = fx.Module(
	"server",
	logger.Module,
	appconfig.Module,
	http.Module,
	validator.Module,
	handlers.Module,
	services.Module,
	repositories.Module,
	models.Module,
	db.Module,
	tasks.Module,
)

func Run() {
	cli.DefaultCommand = "http:run"
	fx.New(
		cli.GetModule(),
		Module,
		fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
			logOption := fxevent.ZapLogger{Logger: logger}
			logOption.UseLogLevel(zapcore.DebugLevel)
			return &logOption
		}),
	).Run()
}
