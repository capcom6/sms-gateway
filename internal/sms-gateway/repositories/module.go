package repositories

import "go.uber.org/fx"

var Module = fx.Module(
	"repositories",
	fx.Provide(
		NewDevicesRepository,
		NewMessagesRepository,
		NewUsersRepository,
	),
)
