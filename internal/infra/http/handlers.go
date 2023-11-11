package http

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type ApiHanlder interface {
	Register(app fiber.Router)
}

type RootHanlder interface {
	Register(app *fiber.App)
}

func AsApiHandler(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(ApiHanlder)),
		fx.ResultTags(`group:"api-routes"`),
	)
}

func AsRootHandler(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(RootHanlder)),
		fx.ResultTags(`group:"root-routes"`),
	)
}

func errorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	// Retrieve the custom status code if it's an fiber.*Error
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	// Send json error
	return c.Status(code).JSON(&fiber.Map{
		"message": err.Error(),
	})
}
