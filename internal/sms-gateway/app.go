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
	"github.com/capcom6/sms-gateway/internal/sms-gateway/modules/auth"
	appdb "github.com/capcom6/sms-gateway/internal/sms-gateway/modules/db"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/modules/devices"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/modules/health"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/modules/messages"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/modules/metrics"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/modules/push"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/modules/webhooks"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/repositories"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Module = fx.Module(
	"server",
	logger.Module,
	appconfig.Module,
	appdb.Module,
	http.Module,
	validator.Module,
	handlers.Module,
	auth.Module,
	push.Module,
	repositories.Module,
	db.Module,
	messages.Module,
	health.Module,
	webhooks.Module,
	devices.Module,
	metrics.Module,
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

	Server          *http.Server
	MessagesService *messages.Service
	PushService     *push.Service
}

func Start(p StartParams) error {
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	p.LC.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			p.MessagesService.RunBackgroundTasks(ctx, wg)

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
