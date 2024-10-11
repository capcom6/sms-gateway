package webhooks

import (
	"github.com/android-sms-gateway/client-go/smsgateway"
)

func webhookToDTO(model *Webhook) smsgateway.Webhook {
	return smsgateway.Webhook{
		ID:    model.ExtID,
		URL:   model.URL,
		Event: model.Event,
	}
}
