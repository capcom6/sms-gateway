package services

import (
	"context"
	"sync"
	"time"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"github.com/capcom6/sms-gateway/pkg/types/cache"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/api/option"
)

type PushServiceParams struct {
	fx.In

	Config PushServiceConfig
	Logger *zap.Logger
}

type PushService struct {
	Config PushServiceConfig

	Logger *zap.Logger

	client *messaging.Client
	mux    sync.Mutex

	cache *cache.Cache[map[string]string]
}

type PushServiceConfig struct {
	CredentialsJSON string
	Timeout         time.Duration
}

func NewPushService(params PushServiceParams) *PushService {
	if params.Config.Timeout == 0 {
		params.Config.Timeout = time.Second
	}

	return &PushService{
		Config: params.Config,
		Logger: params.Logger,
		cache:  cache.New[map[string]string](),
	}
}

// init
func (s *PushService) init(ctx context.Context) (err error) {
	s.mux.Lock()
	defer s.mux.Unlock()

	if s.client != nil {
		return
	}

	opt := option.WithCredentialsJSON([]byte(s.Config.CredentialsJSON))

	var app *firebase.App
	app, err = firebase.NewApp(ctx, nil, opt)

	if err != nil {
		return
	}

	s.client, err = app.Messaging(ctx)

	return
}

func (s *PushService) sendAll(ctx context.Context) {
	if err := s.init(ctx); err != nil {
		s.Logger.Error("Can't init push service", zap.Error(err))
		return
	}

	targets := s.cache.Drain()
	if len(targets) == 0 {
		return
	}

	s.Logger.Info("Sending messages", zap.Int("count", len(targets)))
	for token, data := range targets {
		singleCtx, cancel := context.WithTimeout(ctx, s.Config.Timeout)
		if err := s.sendSingle(singleCtx, token, data); err != nil {
			s.Logger.Error("Can't send message", zap.String("token", token), zap.Error(err))
		}
		cancel()
	}
}

func (s *PushService) sendSingle(ctx context.Context, token string, data map[string]string) error {
	_, err := s.client.Send(ctx, &messaging.Message{
		Data: data,
		Android: &messaging.AndroidConfig{
			Priority: "high",
		},
		Token: token,
	})

	return err
}

func (s *PushService) Run(ctx context.Context) {
	ticker := time.NewTicker(time.Second)
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

func (s *PushService) Enqueue(ctx context.Context, token string, data map[string]string) error {
	s.cache.Set(token, data)

	return nil
}
