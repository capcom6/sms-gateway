package config

import (
	"bitbucket.org/capcom6/smsgatewaybackend/internal/infra/db"
	"bitbucket.org/capcom6/smsgatewaybackend/internal/infra/http"
	"bitbucket.org/capcom6/smsgatewaybackend/internal/smsgateway/services"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"appconfig",
	fx.Provide(
		fx.Annotate(
			func() any {
				return &Config{}
			},
			fx.ResultTags(`name:"config:source"`),
		),
	),
	fx.Provide(
		fx.Annotate(
			func(cfg any) Config {
				return *cfg.(*Config)
			},
			fx.ParamTags(`name:"config:result"`),
		),
	),
	// fx.Provide(GetConfig),
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
