package handlers

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Handler struct {
	Logger    *zap.Logger
	Validator *validator.Validate
}

func (h *Handler) BodyParserValidator(c *fiber.Ctx, out any) error {
	if err := c.BodyParser(out); err != nil {
		return fmt.Errorf("can't parse body: %w", err)
	}

	if h.Validator == nil {
		return nil
	}

	return h.Validator.Struct(out)
}

func (h *Handler) QueryParserValidator(c *fiber.Ctx, out any) error {
	if err := c.QueryParser(out); err != nil {
		return fmt.Errorf("can't parse query: %w", err)
	}

	if h.Validator == nil {
		return nil
	}

	return h.Validator.Struct(out)
}

func (h *Handler) ParamsParserValidator(c *fiber.Ctx, out any) error {
	if err := c.ParamsParser(out); err != nil {
		return fmt.Errorf("can't parse params: %w", err)
	}

	if h.Validator == nil {
		return nil
	}

	return h.Validator.Struct(out)
}
