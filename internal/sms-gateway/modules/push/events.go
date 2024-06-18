package push

import (
	"encoding/json"

	"github.com/android-sms-gateway/client-go/smsgateway"
)

type Event struct {
	Event smsgateway.PushEventType
	Data  any
}

func (e *Event) Map() map[string]string {
	json, _ := json.Marshal(e.Data)

	return map[string]string{
		"event": string(e.Event),
		"data":  string(json),
	}
}

func NewEvent(event smsgateway.PushEventType, data any) *Event {
	return &Event{
		Event: event,
		Data:  data,
	}
}

func NewMessageEnqueuedEvent() *Event {
	return NewEvent(smsgateway.PushMessageEnqueued, nil)
}

func NewWebhooksUpdatedEvent() *Event {
	return NewEvent(smsgateway.PushWebhooksUpdated, nil)
}
