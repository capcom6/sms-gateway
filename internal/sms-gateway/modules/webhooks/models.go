package webhooks

import (
	"github.com/android-sms-gateway/client-go/smsgateway/webhooks"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/models"
	"gorm.io/gorm"
)

type Webhook struct {
	ID     uint64 `json:"-"    gorm:"->;primaryKey;type:BIGINT UNSIGNED;autoIncrement"`
	ExtID  string `json:"id"   gorm:"not null;type:varchar(36);uniqueIndex:unq_webhooks_user_extid,priority:2"`
	UserID string `json:"-"    gorm:"<-:create;not null;type:varchar(32);uniqueIndex:unq_webhooks_user_extid,priority:1"`

	URL   string             `json:"url"   validate:"required,http_url"   gorm:"not null;type:varchar(256)"`
	Event webhooks.EventType `json:"event" gorm:"not null;type:varchar(32)"`

	User models.User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`

	models.TimedModel
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&Webhook{})
}
