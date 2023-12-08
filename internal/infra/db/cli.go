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

func AutoMigrate(params CommandMigrateParams) error {
	err := params.DB.Transaction(func(tx *gorm.DB) error {
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

	params.Logger.Info("Migrations completed")

	return params.Shut.Shutdown()
}
