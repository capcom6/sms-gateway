package domain

import (
	"encoding/json"

	"github.com/android-sms-gateway/client-go/smsgateway"
)

type Event struct {
	Event smsgateway.PushEventType
	Data  map[string]string
}

func (e *Event) Map() map[string]string {
	json, _ := json.Marshal(e.Data)

	return map[string]string{
		"event": string(e.Event),
		"data":  string(json),
	}
}

func NewEvent(event smsgateway.PushEventType, data map[string]string) *Event {
	return &Event{
		Event: event,
		Data:  data,
	}
}
