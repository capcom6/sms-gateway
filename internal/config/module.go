package config

import (
	"github.com/capcom6/sms-gateway/internal/infra/config"
	"github.com/capcom6/sms-gateway/internal/infra/db"
	"github.com/capcom6/sms-gateway/internal/infra/http"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/services"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module(
	"appconfig",
	fx.Provide(
		func(log *zap.Logger) Config {
			if err := config.LoadConfig(&defaultConfig); err != nil {
				log.Error("Error loading config", zap.Error(err))
			}

			return defaultConfig
		},
	),
	fx.Provide(func(cfg Config) http.Config {
		return http.Config{
			Listen: cfg.HTTP.Listen,
		}
	}),
	fx.Provide(func(cfg Config) db.Config {
		return db.Config{
			Host:     cfg.Database.Host,
			Port:     cfg.Database.Port,
			User:     cfg.Database.User,
			Password: cfg.Database.Password,
			Database: cfg.Database.Database,
			Timezone: cfg.Database.Timezone,
		}
	}),
	fx.Provide(func(cfg Config) services.PushServiceConfig {
		return services.PushServiceConfig{
			CredentialsJSON: cfg.FCM.CredentialsJSON,
		}
	}),
)
