package models

import (
	"time"

	shared "github.com/capcom6/sms-gateway/internal/shared/models"
)

type MessageState = shared.MessageState

const (
	MessageStatePending   MessageState = shared.MessageStatePending
	MessageStateProcessed MessageState = shared.MessageStateProcessed
	MessageStateSent      MessageState = shared.MessageStateSent
	MessageStateDelivered MessageState = shared.MessageStateDelivered
	MessageStateFailed    MessageState = shared.MessageStateFailed
)

type User struct {
	ID           string   `gorm:"primaryKey;type:varchar(32)"`
	PasswordHash string   `gorm:"not null;type:varchar(72)"`
	Devices      []Device `gorm:"-,foreignKey:UserID;constraint:OnDelete:CASCADE"`

	shared.TimedMixin
}

type Device struct {
	ID        string  `gorm:"primaryKey;type:char(21)"`
	Name      *string `gorm:"type:varchar(128)"`
	AuthToken string  `gorm:"not null;uniqueIndex;type:char(21)"`
	PushToken *string `gorm:"type:varchar(256)"`

	LastSeen time.Time `gorm:"not null;autocreatetime:false;default:CURRENT_TIMESTAMP(3)"`

	UserID string `gorm:"not null;type:varchar(32)"`

	shared.TimedMixin
}

type Message struct {
	shared.Message

	DeviceID string `gorm:"not null;type:char(21);uniqueIndex:unq_messages_id_device,priority:2;index:idx_messages_device_state"`

	Device Device `gorm:"foreignKey:DeviceID;constraint:OnDelete:CASCADE"`
}

type MessageRecipient = shared.MessageRecipient
