package repositories

import (
	"github.com/capcom6/sms-gateway/internal/sms-gateway/models"
	"gorm.io/gorm"
)

var (
	ErrDeviceNotFound = gorm.ErrRecordNotFound
)

type DevicesRepository struct {
	db *gorm.DB
}

func (r *DevicesRepository) Get(id string) (models.Device, error) {
	device := models.Device{}

	return device, r.db.Where("id = ?", id).Take(&device).Error
}

func (r *DevicesRepository) GetByToken(token string) (models.Device, error) {
	device := models.Device{}

	return device, r.db.Where("auth_token = ?", token).Take(&device).Error
}

func (r *DevicesRepository) Insert(device *models.Device) error {
	return r.db.Create(device).Error
}

func (r *DevicesRepository) UpdateToken(id, token string) error {
	return r.db.Model(&models.Device{}).Where("id", id).Update("push_token", token).Error
}

func NewDevicesRepository(db *gorm.DB) *DevicesRepository {
	return &DevicesRepository{
		db: db,
	}
}
