package http

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Server struct {
	Config Config
	App    *fiber.App
	Logger *zap.Logger
}

func NewServer(config Config, app *fiber.App, logger *zap.Logger) *Server {
	return &Server{
		Config: config,
		App:    app,
		Logger: logger,
	}
}

func (s *Server) Start() error {
	s.Logger.Info("Starting server at " + s.Config.Listen + "...")

	return s.App.Listen(s.Config.Listen)
}

func (s *Server) Stop(ctx context.Context) error {
	defer s.Logger.Info("Server stopped")
	return s.App.ShutdownWithContext(ctx)
}
