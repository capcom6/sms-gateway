package db

import (
	"database/sql"
	"io/fs"

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

	Config Config

	Logger *zap.Logger
	DB     *sql.DB
	Shut   fx.Shutdowner
}

type GooseMigrate struct {
	Config Config
	DB     *sql.DB
	Logger *zap.Logger
	Shut   fx.Shutdowner
}

func NewGooseMigrate(params GooseMigrateParams) *GooseMigrate {
	return &GooseMigrate{
		Config: params.Config,
		Logger: params.Logger,
		DB:     params.DB,
		Shut:   params.Shut,
	}
}

func (c *GooseMigrate) Cmd() string {
	return "db:migrate"
}

func (c *GooseMigrate) Run(args ...string) error {
	cmd := "up"
	if len(args) > 0 {
		cmd = args[0]
	}

	if err := goose.SetDialect(c.Config.Dialect); err != nil {
		return err
	}

	migrationsPath := "migrations/" + c.Config.Dialect

	for _, fs := range gooseStorages {
		goose.SetBaseFS(fs)

		switch cmd {
		case "up":
			if err := goose.Up(c.DB, migrationsPath); err != nil {
				return err
			}
		case "down":
			if err := goose.Down(c.DB, migrationsPath); err != nil {
				return err
			}
		}
	}

	c.Logger.Info("Migrations completed")

	return c.Shut.Shutdown()
}
