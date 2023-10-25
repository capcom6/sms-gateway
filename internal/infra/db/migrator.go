package db

import "gorm.io/gorm"

type Migrator interface {
	Migrate(*gorm.DB) error
}
