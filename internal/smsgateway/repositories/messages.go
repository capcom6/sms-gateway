package repositories

import (
	"bitbucket.org/capcom6/smsgatewaybackend/internal/smsgateway/models"
	"gorm.io/gorm"
)

type MessagesRepository struct {
	db *gorm.DB
}

func (r *MessagesRepository) SelectPending(deviceID string) (messages []models.Message, err error) {
	err = r.db.
		Where("device_id = ? AND state = ?", deviceID, models.MessageStatePending).
		Order("id").
		Preload("Recipients").
		Find(&messages).
		Error

	return
}

func NewMessagesRepository(db *gorm.DB) *MessagesRepository {
	return &MessagesRepository{
		db: db,
	}
}
