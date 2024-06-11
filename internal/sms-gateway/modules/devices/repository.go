package devices

import (
	"errors"
	"time"

	"github.com/capcom6/sms-gateway/internal/sms-gateway/models"
	"gorm.io/gorm"
)

var (
	ErrNotFound      = gorm.ErrRecordNotFound
	ErrInvalidFilter = errors.New("invalid filter")
	ErrMoreThanOne   = errors.New("more than one record")
)

type repository struct {
	db *gorm.DB
}

func (r *repository) Select(filter ...SelectFilter) ([]models.Device, error) {
	if len(filter) == 0 {
		return nil, ErrInvalidFilter
	}

	f := newFilter(filter...)
	devices := []models.Device{}

	return devices, f.apply(r.db).Find(&devices).Error
}

func (r *repository) Get(filter ...SelectFilter) (models.Device, error) {
	devices, err := r.Select(filter...)
	if err != nil {
		return models.Device{}, err
	}

	if len(devices) == 0 {
		return models.Device{}, ErrNotFound
	}

	if len(devices) > 1 {
		return models.Device{}, ErrMoreThanOne
	}

	return devices[0], nil
}

func (r *repository) Insert(device *models.Device) error {
	return r.db.Create(device).Error
}

func (r *repository) UpdateToken(id, token string) error {
	return r.db.Model(&models.Device{}).Where("id", id).Update("push_token", token).Error
}

func (r *repository) UpdateLastSeen(id string) error {
	return r.db.Model(&models.Device{}).Where("id", id).Update("last_seen", time.Now()).Error
}

func newDevicesRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}
