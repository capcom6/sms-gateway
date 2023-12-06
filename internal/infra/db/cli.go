package db

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CommandMigrateParams struct {
	fx.In

	Logger *zap.Logger
	DB     *gorm.DB
	Shut   fx.Shutdowner
}

type CommandMigrate struct {
	Logger *zap.Logger
	DB     *gorm.DB
	Shut   fx.Shutdowner
}

func NewCommandMigrate(params CommandMigrateParams) *CommandMigrate {
	return &CommandMigrate{
		Logger: params.Logger,
		DB:     params.DB,
		Shut:   params.Shut,
	}
}

func (c *CommandMigrate) Cmd() string {
	return "db:auto-migrate"
}

func (c *CommandMigrate) Run(args ...string) error {
	err := c.DB.Transaction(func(tx *gorm.DB) error {
		for _, v := range migrations {
			if err := v(tx); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	c.Logger.Info("Migrations completed")

	return c.Shut.Shutdown()
}
