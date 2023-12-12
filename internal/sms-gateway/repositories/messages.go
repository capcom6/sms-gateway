package repositories

import (
	"github.com/capcom6/sms-gateway/internal/sms-gateway/models"
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

func (r *MessagesRepository) Get(ID string, filter MessagesSelectFilter, options ...MessagesSelectOptions) (message models.Message, err error) {
	query := r.db.Model(&message).
		Where("ext_id = ?", ID)

	if filter.DeviceID != "" {
		query = query.Where("device_id = ?", filter.DeviceID)
	}

	if len(options) > 0 {
		if options[0].WithRecipients {
			query = query.Preload("Recipients")
		}
		if options[0].WithDevice {
			query = query.Preload("Device")
		}
	}

	err = query.Take(&message).Error

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
			if err := tx.Model(&v).Where("message_id = ?", message.ID).Select("State", "Error").Updates(&v).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func (r *MessagesRepository) HashProcessed() error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&models.MessageRecipient{}).
			Where("message_id IN (?)", tx.Model(&models.Message{}).Select("id").Where("is_hashed = ? AND state <> ?", false, models.MessageStatePending)).
			Update("phone_number", gorm.Expr("LEFT(SHA2(phone_number, 256), 16)")).
			Error
		if err != nil {
			return err
		}

		return tx.Model(&models.Message{}).
			Where("is_hashed = ? AND state <> ?", false, models.MessageStatePending).
			Updates(map[string]interface{}{"is_hashed": true, "message": gorm.Expr("SHA2(message, 256)")}).
			Error
	})
}

func NewMessagesRepository(db *gorm.DB) *MessagesRepository {
	return &MessagesRepository{
		db: db,
	}
}

// /////////////////////////////////////////////////////////////////////////////
type MessagesSelectFilter struct {
	DeviceID string
}

type MessagesSelectOptions struct {
	WithRecipients bool
	WithDevice     bool
}
