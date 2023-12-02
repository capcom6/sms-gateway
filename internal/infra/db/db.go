package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"moul.io/zapgorm2"
)

func New(params Params) (*gorm.DB, error) {
	cfgGorm := makeConfig(params)

	return gorm.Open(mysql.New(mysql.Config{Conn: params.SQL}), cfgGorm)
}

func makeConfig(params Params) *gorm.Config {
	log := zapgorm2.New(params.Logger)
	log.LogLevel = logger.Info
	log.SetAsDefault()

	return &gorm.Config{
		Logger: log,
	}
}
