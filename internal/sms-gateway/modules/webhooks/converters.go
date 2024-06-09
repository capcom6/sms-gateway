package webhooks

import "github.com/capcom6/sms-gateway/pkg/smsgateway"

func webhookToDTO(model *Webhook) smsgateway.WebhookDTO {
	return smsgateway.WebhookDTO{
		ID:    model.ExtID,
		URL:   model.URL,
		Event: model.Event,
	}
}
