package webhooks

import (
	"fmt"

	"github.com/capcom6/sms-gateway/internal/sms-gateway/handlers/base"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/models"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/modules/auth"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type handlerParams struct {
	fx.In

	WebhooksSvc *Service

	Validator *validator.Validate
	Logger    *zap.Logger
}

type Handler struct {
	base.Handler

	webhooksSvc *Service

	logger *zap.Logger
}

func (h *Handler) get(user models.User, c *fiber.Ctx) error {
	items, err := h.webhooksSvc.Select(user.ID)
	if err != nil {
		return fmt.Errorf("can't select webhooks: %w", err)
	}

	return c.JSON(items)
}

func (h *Handler) post(user models.User, c *fiber.Ctx) error {
	dto := WebhookDTO{}

	if err := h.BodyParserValidator(c, &dto); err != nil {
		return err
	}

	if err := h.webhooksSvc.Replace(user.ID, dto); err != nil {
		return fmt.Errorf("can't write webhook: %w", err)
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (h *Handler) delete(user models.User, c *fiber.Ctx) error {
	id := c.Params("id")

	if err := h.webhooksSvc.Delete(user.ID, WithExtID(id)); err != nil {
		return fmt.Errorf("can't delete webhook: %w", err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *Handler) Register(router fiber.Router) {
	router.Get("/", auth.WithUser(h.get))
	router.Post("/", auth.WithUser(h.post))
	router.Delete("/:id", auth.WithUser(h.delete))
}

func NewHandler(params handlerParams) *Handler {
	return &Handler{
		Handler: base.Handler{
			Logger:    params.Logger.Named("webhooks"),
			Validator: params.Validator,
		},
		webhooksSvc: params.WebhooksSvc,
		logger:      params.Logger,
	}
}
