package push

import (
	"context"
	"time"

	"github.com/capcom6/sms-gateway/pkg/types/cache"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Config struct {
	Mode Mode

	ClientOptions map[string]string

	Debounce time.Duration
	Timeout  time.Duration
}

type Params struct {
	fx.In

	Config Config

	Client client

	Logger *zap.Logger
}

type Service struct {
	config Config

	client client

	cache *cache.Cache[map[string]string]

	logger *zap.Logger
}

func New(params Params) *Service {
	if params.Config.Timeout == 0 {
		params.Config.Timeout = time.Second
	}
	if params.Config.Debounce < 5*time.Second {
		params.Config.Debounce = 5 * time.Second
	}

	return &Service{
		config: params.Config,
		client: params.Client,
		cache:  cache.New[map[string]string](),
		logger: params.Logger,
	}
}

// Run runs the service with the provided context if a debounce is set.
func (s *Service) Run(ctx context.Context) {
	ticker := time.NewTicker(s.config.Debounce)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			s.sendAll(ctx)
		}
	}
}

// Enqueue adds the data to the cache and immediately sends all messages if the debounce is 0.
func (s *Service) Enqueue(ctx context.Context, token string, event *Event) error {
	s.cache.Set(token, event.Map())

	return nil
}

// sendAll sends messages to all targets from the cache after initializing the service.
func (s *Service) sendAll(ctx context.Context) {
	targets := s.cache.Drain()
	if len(targets) == 0 {
		return
	}

	s.logger.Info("Sending messages", zap.Int("count", len(targets)))
	ctx, cancel := context.WithTimeout(ctx, s.config.Timeout)
	if err := s.client.Send(ctx, targets); err != nil {
		s.logger.Error("Can't send messages", zap.Error(err))
	}
	cancel()
}
