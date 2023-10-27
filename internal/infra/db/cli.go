package db

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CommandMigrateParams struct {
	fx.In

	Logger     *zap.Logger
	DB         *gorm.DB
	Migrations []Migrator `group:"migrations"`
	Shut       fx.Shutdowner
}

type CommandMigrate struct {
	Logger     *zap.Logger
	DB         *gorm.DB
	Migrations []Migrator `group:"migrations"`
	Shut       fx.Shutdowner
}

func NewCommandMigrate(params CommandMigrateParams) *CommandMigrate {
	return &CommandMigrate{
		Logger:     params.Logger,
		DB:         params.DB,
		Migrations: params.Migrations,
		Shut:       params.Shut,
	}
}

func (c *CommandMigrate) Cmd() string {
	return "migrate"
}

func (c *CommandMigrate) Run(args ...string) error {
	for _, v := range c.Migrations {
		if err := v.Migrate(c.DB); err != nil {
			return err
		}
	}
	c.Logger.Info("Migrations completed")

	return c.Shut.Shutdown()
}
