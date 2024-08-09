package webhooks

import (
	"github.com/android-sms-gateway/client-go/smsgateway/webhooks"
)

func webhookToDTO(model *Webhook) webhooks.Webhook {
	return webhooks.Webhook{
		ID:    model.ExtID,
		URL:   model.URL,
		Event: model.Event,
	}
}
