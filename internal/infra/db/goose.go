package db

import (
	"database/sql"
	"io/fs"

	"github.com/capcom6/sms-gateway/internal/infra/cli"
	"github.com/pressly/goose/v3"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type GooseStorage fs.FS

var gooseStorages = []GooseStorage{}

func RegisterGoose(storage GooseStorage) {
	gooseStorages = append(gooseStorages, storage)
}

type GooseMigrateParams struct {
	fx.In

	Args cli.Args

	Config Config

	Logger *zap.Logger
	DB     *sql.DB
	Shut   fx.Shutdowner
}

func Migrate(params GooseMigrateParams) error {
	cmd := "up"
	if len(params.Args) > 0 {
		cmd = params.Args[0]
	}

	if err := goose.SetDialect(params.Config.Dialect); err != nil {
		return err
	}

	migrationsPath := "migrations/" + params.Config.Dialect

	for _, fs := range gooseStorages {
		goose.SetBaseFS(fs)

		switch cmd {
		case "up":
			if err := goose.Up(params.DB, migrationsPath); err != nil {
				return err
			}
		case "down":
			if err := goose.Down(params.DB, migrationsPath); err != nil {
				return err
			}
		}
	}

	params.Logger.Info("Migrations completed")

	return params.Shut.Shutdown()
}
