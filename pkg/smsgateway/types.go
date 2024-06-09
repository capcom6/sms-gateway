package smsgateway

type WebhookEvent string

const (
	WebhookEventSmsReceived WebhookEvent = "sms:received"
)
