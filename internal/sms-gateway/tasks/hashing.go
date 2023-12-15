package tasks

import (
	"context"
	"time"

	"github.com/capcom6/sms-gateway/internal/sms-gateway/services"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type HashingTaskConfig struct {
	Interval time.Duration
}

type HashingTaskParams struct {
	fx.In

	MessagesSvc *services.MessagesService
	Config      HashingTaskConfig
	Logger      *zap.Logger
}

type HashingTask struct {
	MessagesSvc *services.MessagesService
	Config      HashingTaskConfig
	Logger      *zap.Logger
}

func (t *HashingTask) Run(ctx context.Context) {
	t.Logger.Info("Starting hashing task...")
	for {
		select {
		case <-ctx.Done():
			t.Logger.Info("Stopping hashing task...")
			return
		case <-time.After(t.Config.Interval):
			t.Logger.Debug("Hashing messages...")
			if err := t.MessagesSvc.HashProcessed(); err != nil {
				t.Logger.Error("Failed to hash processed messages", zap.Error(err))
			}
		}
	}
}

func NewHashingTask(params HashingTaskParams) *HashingTask {
	return &HashingTask{
		MessagesSvc: params.MessagesSvc,
		Config:      params.Config,
		Logger:      params.Logger,
	}
}
