package config

import (
	"go.uber.org/fx"
)

var Module = fx.Module(
	"config",
	fx.Provide(
		fx.Annotate(
			New,
			fx.ResultTags(`name:"config:result"`),
		),
	),
)

// func MakeAppModule(defaults *Config) fx.Option {
// 	return fx.Module(
// 		"appconfig",
// 		fx.Provide(
// 			fx.Annotate(
// 				func() any {
// 					return defaults
// 				},
// 				fx.ResultTags(`name:"config:source"`),
// 			),
// 		),
// 		fx.Provide(
// 			fx.Annotate(
// 				func(cfg any) Config {
// 					fmt.Printf("%+#v", cfg)
// 					return *cfg.(*Config)
// 				},
// 				fx.ParamTags(`name:"config:result"`),
// 			),
// 		),
// 	)
// }
