package http

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type ApiHanlder interface {
	Register(app fiber.Router)
}

func AsApiHandler(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(ApiHanlder)),
		fx.ResultTags(`group:"api-routes"`),
	)
}
