package webhooks

func webhookToDTO(model *Webhook) WebhookDTO {
	return WebhookDTO{
		ID:    model.ExtID,
		URL:   model.URL,
		Event: model.Event,
	}
}
