package models

import (
	"embed"

	"github.com/capcom6/sms-gateway/internal/infra/db"
)

//go:embed migrations
var migrations embed.FS

func GetGooseStorage() db.GooseStorage {
	return db.GooseStorage{
		FS: migrations,
	}
}
