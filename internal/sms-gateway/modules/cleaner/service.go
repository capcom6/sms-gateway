package cleaner

import (
	"context"
	"time"

	"go.uber.org/zap"
)

type Service struct {
	targets []Cleanable

	logger *zap.Logger
}

func New(targets []Cleanable, logger *zap.Logger) *Service {
	return &Service{
		targets: targets,
		logger:  logger,
	}
}

func (s *Service) Run(ctx context.Context) {
	ticker := time.NewTicker(24 * time.Hour)
	defer ticker.Stop()

	s.logger.Info("Cleaner started")
	defer s.logger.Info("Cleaner stopped")
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			s.clean(ctx)
		}
	}
}

func (s *Service) clean(ctx context.Context) {
	s.logger.Info("Cleaning...")
	defer s.logger.Info("Cleaning...Done")

	for _, target := range s.targets {
		select {
		case <-ctx.Done():
			return
		default:
			if err := target.Clean(ctx); err != nil {
				s.logger.Error("Can't clean target", zap.Error(err))
			}
		}
	}
}
