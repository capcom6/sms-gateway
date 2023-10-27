package http

import (
	"time"

	"bitbucket.org/capcom6/smsgatewaybackend/internal/infra/http/jsonify"
	"bitbucket.org/capcom6/smsgatewaybackend/internal/infra/http/statuscode"
	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

const (
	ReadTimeout  = 5 * time.Second
	WriteTimeout = 5 * time.Second
	IdleTimeout  = 60 * time.Second
)

func New(params Params) (*fiber.App, error) {
	app := fiber.New(fiber.Config{
		ReadTimeout:  ReadTimeout,
		WriteTimeout: WriteTimeout,
		IdleTimeout:  IdleTimeout,
	})

	app.Use(recover.New())
	app.Use(fiberzap.New(fiberzap.Config{
		Logger: params.Logger,
	}))

	api := app.Group("/api")
	api.Use(cors.New())
	api.Use(jsonify.New())
	for _, handler := range params.ApiHandlers {
		handler.Register(api)
	}

	app.Use(statuscode.New())

	// params.LC.Append(fx.Hook{
	// 	OnStart: func(ctx context.Context) error {
	// 		go func() {
	// 			err := app.Listen(config.Listen)
	// 			if err != nil {
	// 				params.Logger.Error("Error starting server", zap.Error(err))
	// 			}
	// 		}()

	// 		return nil
	// 	},
	// 	OnStop: func(ctx context.Context) error {
	// 		return app.ShutdownWithContext(ctx)
	// 	},
	// })

	return app, nil
}
