package repositories

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module(
	"repositories",
	fx.Decorate(func(log *zap.Logger) *zap.Logger {
		return log.Named("repositories")
	}),
	fx.Provide(
		NewMessagesRepository,
		NewUsersRepository,
	),
)
