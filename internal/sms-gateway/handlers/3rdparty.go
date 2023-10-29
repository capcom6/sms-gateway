package handlers

import (
	"errors"

	"github.com/capcom6/sms-gateway/internal/sms-gateway/models"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/repositories"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/services"
	"github.com/capcom6/sms-gateway/pkg/smsgateway"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"go.uber.org/zap"
)

type thirdPartyHandler struct {
	Handler

	authSvc     *services.AuthService
	messagesSvc *services.MessagesService
}

//	@Summary		Поставить сообщение в очередь
//	@Description	Ставит сообщение в очередь на отправку. Если идентификатор не указан, то он будет сгенерирован автоматически
//	@Security		ApiAuth
//	@Tags			Пользователь, Сообщения
//	@Accept			json
//	@Produce		json
//	@Param			request	body		smsgateway.Message			true	"Сообщение"
//	@Success		201		{object}	smsgateway.MessageState		"Сообщение поставлено в очередь"
//	@Failure		401		{object}	smsgateway.ErrorResponse	"Ошибка авторизации"
//	@Failure		400		{object}	smsgateway.ErrorResponse	"Некорректный запрос"
//	@Failure		500		{object}	smsgateway.ErrorResponse	"Внутренняя ошибка сервера"
//	@Router			/3rdparty/v1/message [post]
//
// Поставить сообщение в очередь
func (h *thirdPartyHandler) postMessage(user models.User, c *fiber.Ctx) error {
	req := smsgateway.Message{}
	if err := h.BodyParserValidator(c, &req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if len(user.Devices) < 1 {
		return fiber.NewError(fiber.StatusBadRequest, "Нет ни одного устройства в учетной записи")
	}

	device := user.Devices[0]
	state, err := h.messagesSvc.Enqeue(device, req)
	if err != nil {
		if errors.Is(err, services.ErrValidation) {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		return err
	}

	return c.Status(fiber.StatusCreated).JSON(state)
}

//	@Summary		Получить состояние сообщения
//	@Description	Возвращает состояние сообщения по его ID
//	@Security		ApiAuth
//	@Tags			Пользователь, Сообщения
//	@Produce		json
//	@Param			id	path		string						true	"ИД сообщения"
//	@Success		200	{object}	smsgateway.MessageState		"Состояние сообщения"
//	@Failure		401	{object}	smsgateway.ErrorResponse	"Ошибка авторизации"
//	@Failure		400	{object}	smsgateway.ErrorResponse	"Некорректный запрос"
//	@Failure		500	{object}	smsgateway.ErrorResponse	"Внутренняя ошибка сервера"
//	@Router			/3rdparty/v1/message [get]
//
// Получить состояние сообщения
func (h *thirdPartyHandler) getMessage(user models.User, c *fiber.Ctx) error {
	id := c.Params("id")

	state, err := h.messagesSvc.GetState(user, id)
	if err != nil {
		if errors.Is(err, repositories.ErrMessageNotFound) {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}

		return err
	}

	return c.JSON(state)
}

func (h *thirdPartyHandler) authorize(handler func(models.User, *fiber.Ctx) error) fiber.Handler {
	return func(c *fiber.Ctx) error {
		username := c.Locals("username").(string)
		password := c.Locals("password").(string)

		user, err := h.authSvc.AuthorizeUser(username, password)
		if err != nil {
			h.Logger.Error("failed to authorize user", zap.Error(err))
			return fiber.ErrUnauthorized
		}

		return handler(user, c)
	}
}

func (h *thirdPartyHandler) Register(router fiber.Router) {
	router = router.Group("/3rdparty/v1")

	router.Use(basicauth.New(basicauth.Config{
		Authorizer: func(username string, password string) bool {
			return len(username) > 0 && len(password) > 0
		},
	}))

	router.Post("/message", h.authorize(h.postMessage))
	router.Get("/message/:id", h.authorize(h.getMessage))
}

func newThirdPartyHandler(logger *zap.Logger, validator *validator.Validate, authSvc *services.AuthService, messagesSvc *services.MessagesService) *thirdPartyHandler {
	return &thirdPartyHandler{
		Handler: Handler{
			Logger:    logger,
			Validator: validator,
		},
		authSvc:     authSvc,
		messagesSvc: messagesSvc,
	}
}
