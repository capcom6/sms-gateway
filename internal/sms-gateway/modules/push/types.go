package push

import (
	"context"

	"github.com/android-sms-gateway/client-go/smsgateway"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/modules/push/domain"
)

type Mode string
type Event = domain.Event

const (
	ModeFCM      Mode = "fcm"
	ModeUpstream Mode = "upstream"
)

type client interface {
	Open(ctx context.Context) error
	Send(ctx context.Context, messages map[string]Event) error
	Close(ctx context.Context) error
}

func NewMessageEnqueuedEvent() *Event {
	return domain.NewEvent(smsgateway.PushMessageEnqueued, nil)
}

func NewWebhooksUpdatedEvent() *Event {
	return domain.NewEvent(smsgateway.PushWebhooksUpdated, nil)
}
