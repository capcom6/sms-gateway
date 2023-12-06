package db

import (
	"gorm.io/gorm"
)

type Migrator func(*gorm.DB) error

var migrations = []Migrator{}

func RegisterMigration(migrator Migrator) {
	migrations = append(migrations, migrator)
}
