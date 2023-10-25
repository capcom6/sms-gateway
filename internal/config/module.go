package config

import (
	"bitbucket.org/capcom6/smsgatewaybackend/internal/infra/db"
	"bitbucket.org/capcom6/smsgatewaybackend/internal/infra/http"
	"bitbucket.org/capcom6/smsgatewaybackend/internal/smsgateway/services"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"config",
	fx.Provide(GetConfig),
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
		}
	}),
	fx.Provide(func(cfg Config) services.PushServiceConfig {
		return services.PushServiceConfig{
			CredentialsJSON: cfg.FCM.CredentialsJSON,
		}
	}),
)
