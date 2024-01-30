package smsgateway

import (
	"context"
	"sync"

	"github.com/capcom6/go-infra-fx/cli"
	"github.com/capcom6/go-infra-fx/db"
	"github.com/capcom6/go-infra-fx/http"
	"github.com/capcom6/go-infra-fx/logger"
	"github.com/capcom6/go-infra-fx/validator"
	appconfig "github.com/capcom6/sms-gateway/internal/config"
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
	cli.DefaultCommand = "start"
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

type StartParams struct {
	fx.In

	LC     fx.Lifecycle
	Logger *zap.Logger
	Shut   fx.Shutdowner

	Server      *http.Server
	HashingTask *tasks.HashingTask
	PushService *services.PushService
}

func Start(p StartParams) error {
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	p.LC.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			wg.Add(1)
			go func() {
				defer wg.Done()
				p.HashingTask.Run(ctx)
			}()

			wg.Add(1)
			go func() {
				defer wg.Done()
				p.PushService.Run(ctx)
			}()

			wg.Add(1)
			go func() {
				defer wg.Done()
				if err := p.Server.Start(); err != nil {
					p.Logger.Error("Error starting server", zap.Error(err))
					_ = p.Shut.Shutdown()
				}
			}()

			p.Logger.Info("Service started")

			return nil
		},
		OnStop: func(_ context.Context) error {
			cancel()
			_ = p.Server.Stop(ctx)
			wg.Wait()

			p.Logger.Info("Service stopped")

			return nil
		},
	})

	return nil
}

func init() {
	cli.Register("start", Start)
}
