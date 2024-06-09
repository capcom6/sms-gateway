package smsgateway

type WebhookDTO struct {
	ID    string       `json:"id"    validate:"max=36"            example:"123e4567-e89b-12d3-a456-426614174000"`
	URL   string       `json:"url"   validate:"required,http_url" example:"https://example.com/webhook"`
	Event WebhookEvent `json:"event" validate:"required"          example:"sms:received"`
}
