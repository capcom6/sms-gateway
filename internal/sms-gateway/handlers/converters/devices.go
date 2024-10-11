package converters

import (
	"github.com/android-sms-gateway/client-go/smsgateway"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/models"
)

func DeviceToDTO(device *models.Device) *smsgateway.Device {
	if device.IsEmpty() {
		return nil
	}

	return &smsgateway.Device{
		ID:        device.ID,
		Name:      *device.Name,
		CreatedAt: device.CreatedAt,
		UpdatedAt: device.UpdatedAt,
		DeletedAt: device.DeletedAt,
		LastSeen:  device.LastSeen,
	}
}
