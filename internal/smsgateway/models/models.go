package models

import (
	"time"

	"gorm.io/gorm"
)

type MessageState string

const (
	MessageStatePending   MessageState = "Pending"
	MessageStateProcessed MessageState = "Processes"
	MessageStateSent      MessageState = "Sent"
	MessageStateDelivered MessageState = "Delivered"
	MessageStateFailed    MessageState = "Failed"
)

type TimedModel struct {
	CreatedAt time.Time `gorm:"not null;autocreatetime:false;default:CURRENT_TIMESTAMP(3)"`
	UpdatedAt time.Time `gorm:"not null;autoupdatetime:false;default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3)"`
	DeletedAt gorm.DeletedAt
}

type User struct {
	ID           string   `gorm:"primaryKey;type:varchar(32)"`
	PasswordHash string   `gorm:"not null;type:varchar(72)"`
	Devices      []Device `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`

	TimedModel
}

type Device struct {
	ID        string  `gorm:"primaryKey;type:char(21)"`
	Name      *string `gorm:"type:varchar(128)"`
	AuthToken string  `gorm:"not null;uniqueIndex;type:char(21)"`
	PushToken *string `gorm:"type:varchar(256)"`

	UserID string `gorm:"not null;type:varchar(32)"`

	TimedModel
}

type Message struct {
	ID       uint64       `gorm:"primaryKey;type:BIGINT UNSIGNED;autoIncrement"`
	DeviceID string       `gorm:"not null;type:char(21);uniqueIndex:unq_messages_id_device,priority:2;index:idx_messages_device_state"`
	ExtID    string       `gorm:"not null;type:varchar(36);uniqueIndex:unq_messages_id_device,priority:1"`
	Message  string       `gorm:"not null;type:tinytext"`
	State    MessageState `gorm:"not null;type:enum('Pending','Sent','Processed','Delivered','Failed');default:Pending;index:idx_messages_device_state"`

	Device     Device             `gorm:"foreignKey:DeviceID;constraint:OnDelete:CASCADE"`
	Recipients []MessageRecipient `gorm:"foreignKey:MessageID;constraint:OnDelete:CASCADE"`

	TimedModel
}

type MessageRecipient struct {
	MessageID   uint64       `gorm:"primaryKey;type:BIGINT UNSIGNED"`
	PhoneNumber string       `gorm:"primaryKey;type:varchar(16)"`
	State       MessageState `gorm:"not null;type:enum('Pending','Sent','Processed','Delivered','Failed');default:Pending"`
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&User{}, &Device{}, &Message{}, &MessageRecipient{})
}
