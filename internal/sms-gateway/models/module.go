package models

import (
	"github.com/capcom6/go-infra-fx/db"
)

func init() {
	db.RegisterMigration(Migrate)
	db.RegisterGoose(migrations)
}
