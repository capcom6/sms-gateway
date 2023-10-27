package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"moul.io/zapgorm2"
)

func New(params Params) (*gorm.DB, error) {
	dsn := makeDSN(params.Config)
	cfgGorm := makeConfig(params)

	return gorm.Open(mysql.Open(dsn), cfgGorm)
}

func makeConfig(params Params) *gorm.Config {
	log := zapgorm2.New(params.Logger)
	log.LogLevel = logger.Info
	log.SetAsDefault()

	return &gorm.Config{
		Logger: log,
	}
}

func makeDSN(cfg Config) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4,utf8&parseTime=true&tls=preferred",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database,
	)
}
