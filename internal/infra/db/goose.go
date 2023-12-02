package db

import (
	"database/sql"
	"io/fs"

	"github.com/pressly/goose/v3"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type GooseStorage struct {
	FS fs.FS
}

type GooseMigrateParams struct {
	fx.In

	Config  Config
	Storage GooseStorage

	Logger *zap.Logger
	DB     *sql.DB
	Shut   fx.Shutdowner
}

type GooseMigrate struct {
	Config  Config
	Storage GooseStorage
	DB      *sql.DB
	Logger  *zap.Logger
	Shut    fx.Shutdowner
}

func NewGooseMigrate(params GooseMigrateParams) *GooseMigrate {
	return &GooseMigrate{
		Config:  params.Config,
		Logger:  params.Logger,
		DB:      params.DB,
		Storage: params.Storage,
		Shut:    params.Shut,
	}
}

func (c *GooseMigrate) Cmd() string {
	return "db:migrate"
}

func (c *GooseMigrate) Run(args ...string) error {
	goose.SetBaseFS(c.Storage.FS)

	cmd := "up"
	if len(args) > 0 {
		cmd = args[0]
	}

	if err := goose.SetDialect(c.Config.Dialect); err != nil {
		return err
	}

	migrationsPath := "migrations/" + c.Config.Dialect

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

	c.Logger.Info("Migrations completed")

	return c.Shut.Shutdown()
}
