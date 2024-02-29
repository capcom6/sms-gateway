package shared

import (
	"time"
)

type MessageState string

const (
	MessageStatePending   MessageState = "Pending"
	MessageStateProcessed MessageState = "Processed"
	MessageStateSent      MessageState = "Sent"
	MessageStateDelivered MessageState = "Delivered"
	MessageStateFailed    MessageState = "Failed"
)

type TimedMixin struct {
	CreatedAt time.Time `gorm:"not null;autocreatetime:false;default:CURRENT_TIMESTAMP(3)"`
	UpdatedAt time.Time `gorm:"not null;autoupdatetime:false;default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3)"`
	DeletedAt *time.Time
}

type Message struct {
	ID                 uint64       `gorm:"primaryKey;type:BIGINT UNSIGNED;autoIncrement"`
	ExtID              string       `gorm:"not null;type:varchar(36);uniqueIndex:unq_messages_id_device,priority:1"`
	Text               string       `gorm:"column:message;not null;type:text"`
	State              MessageState `gorm:"not null;type:enum('Pending','Sent','Processed','Delivered','Failed');default:Pending;index:idx_messages_device_state"`
	ValidUntil         *time.Time   `gorm:"type:datetime"`
	SimNumber          *uint8       `gorm:"type:tinyint(1) unsigned"`
	WithDeliveryReport bool         `gorm:"not null;type:tinyint(1) unsigned"`

	IsHashed    bool `gorm:"not null;type:tinyint(1) unsigned;default:0"`
	IsEncrypted bool `gorm:"not null;type:tinyint(1) unsigned;default:0"`

	Recipients []MessageRecipient `gorm:"foreignKey:MessageID;constraint:OnDelete:CASCADE"`

	TimedMixin
}

type MessageRecipient struct {
	MessageID   uint64       `gorm:"primaryKey;type:BIGINT UNSIGNED"`
	PhoneNumber string       `gorm:"primaryKey;type:varchar(128)"`
	State       MessageState `gorm:"not null;type:enum('Pending','Sent','Processed','Delivered','Failed');default:Pending"`
	Error       *string      `gorm:"type:varchar(256)"`
}
