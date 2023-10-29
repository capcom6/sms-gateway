package db

import (
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type Migrator interface {
	Migrate(*gorm.DB) error
}

func AsMigration(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(Migrator)),
		fx.ResultTags(`group:"migrations"`),
	)
}
