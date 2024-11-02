package cleaner

import "context"

type Cleanable interface {
	Clean(ctx context.Context) error
}
