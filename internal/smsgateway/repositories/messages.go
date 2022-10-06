package repositories

import (
	"bitbucket.org/capcom6/smsgatewaybackend/internal/smsgateway/models"
	"gorm.io/gorm"
)

var (
	ErrMessageNotFound = gorm.ErrRecordNotFound
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

func (r *MessagesRepository) Get(deviceID, ID string) (message models.Message, err error) {
	err = r.db.Where("device_id = ? AND ext_id = ?", deviceID, ID).Take(&message).Error

	return
}

func (r *MessagesRepository) Insert(message *models.Message) error {
	return r.db.Create(message).Error
}

func (r *MessagesRepository) UpdateState(message *models.Message) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(message).Select("State").Updates(message).Error; err != nil {
			return err
		}

		for _, v := range message.Recipients {
			if err := tx.Model(&v).Where("message_id = ? AND phone_number = ?", message.ID, v.PhoneNumber).Update("state", v.State).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func NewMessagesRepository(db *gorm.DB) *MessagesRepository {
	return &MessagesRepository{
		db: db,
	}
}
