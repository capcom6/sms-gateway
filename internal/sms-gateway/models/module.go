package models

import (
	"github.com/capcom6/go-infra-fx/db"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"models",
)

func init() {
	db.RegisterMigration(Migrate)
	db.RegisterGoose(migrations)
}
