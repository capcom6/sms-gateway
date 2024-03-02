package push

import "context"

type Mode string

const (
	ModeFCM      Mode = "fcm"
	ModeUpstream Mode = "upstream"
)

type client interface {
	Open(ctx context.Context) error
	Send(ctx context.Context, address string, payload map[string]string) error
	Close(ctx context.Context) error
}
