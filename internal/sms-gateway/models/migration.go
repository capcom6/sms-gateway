package models

import "gorm.io/gorm"

type Migration struct {
}

func (m Migration) Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&User{}, &Device{}, &Message{}, &MessageRecipient{})
}

func NewMigration() *Migration {
	return &Migration{}
}
