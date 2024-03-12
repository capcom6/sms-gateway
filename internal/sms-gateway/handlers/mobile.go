package handlers

import (
	"errors"
	"fmt"
	"strings"

	"github.com/capcom6/go-infra-fx/http/apikey"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/models"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/modules/auth"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/repositories"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/services"
	"github.com/capcom6/sms-gateway/pkg/smsgateway"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/jaevor/go-nanoid"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type mobileHandler struct {
	Handler

	authSvc     *auth.Service
	messagesSvc *services.MessagesService

	idGen func() string
}

//	@Summary		Регистрация устройства
//	@Description	Регистрирует устройство на сервере, генерируя авторизационные данные
//	@Tags			Устройство, Регистрация
//	@Accept			json
//	@Produce		json
//	@Param			request	body		smsgateway.MobileRegisterRequest	true	"Запрос на регистрацию"
//	@Success		201		{object}	smsgateway.MobileRegisterResponse	"Успешная регистрация"
//	@Failure		400		{object}	smsgateway.ErrorResponse			"Некорректный запрос"
//	@Failure		500		{object}	smsgateway.ErrorResponse			"Внутренняя ошибка сервера"
//	@Router			/mobile/v1/device [post]
//
// Регистрация устройства
func (h *mobileHandler) postDevice(c *fiber.Ctx) error {
	req := smsgateway.MobileRegisterRequest{}

	if err := h.BodyParserValidator(c, &req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	id := h.idGen()
	login := strings.ToUpper(id[:6])
	password := strings.ToLower(id[7:])

	user, err := h.authSvc.RegisterUser(login, password)
	if err != nil {
		return fmt.Errorf("can't create user: %w", err)
	}

	device, err := h.authSvc.RegisterDevice(user.ID, req.Name, req.PushToken)
	if err != nil {
		return fmt.Errorf("can't register device: %w", err)
	}

	return c.Status(fiber.StatusCreated).JSON(smsgateway.MobileRegisterResponse{
		Id:       device.ID,
		Token:    device.AuthToken,
		Login:    login,
		Password: password,
	})
}

//	@Summary		Обновление устройства
//	@Description	Обновляет push-токен устройства
//	@Security		MobileToken
//	@Tags			Устройство
//	@Accept			json
//	@Param			request	body	smsgateway.MobileUpdateRequest	true	"Запрос на обновление"
//	@Success		204		"Успешное обновление"
//	@Failure		400		{object}	smsgateway.ErrorResponse	"Некорректный запрос"
//	@Failure		403		{object}	smsgateway.ErrorResponse	"Операция запрещена"
//	@Failure		500		{object}	smsgateway.ErrorResponse	"Внутренняя ошибка сервера"
//	@Router			/mobile/v1/device [patch]
//
// Обновление устройства
func (h *mobileHandler) patchDevice(device models.Device, c *fiber.Ctx) error {
	req := smsgateway.MobileUpdateRequest{}

	if err := h.BodyParserValidator(c, &req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if req.Id != device.ID {
		return fiber.ErrForbidden
	}

	if err := h.authSvc.UpdateDevice(req.Id, req.PushToken); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}

//	@Summary		Получить сообщения для отправки
//	@Description	Возвращает список сообщений, требующих отправки
//	@Security		MobileToken
//	@Tags			Устройство, Сообщения
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		smsgateway.Message			"Список сообщений"
//	@Failure		500	{object}	smsgateway.ErrorResponse	"Внутренняя ошибка сервера"
//	@Router			/mobile/v1/message [get]
//
// Получить сообщения для отправки
func (h *mobileHandler) getMessage(device models.Device, c *fiber.Ctx) error {
	messages, err := h.messagesSvc.SelectPending(device.ID)
	if err != nil {
		return fmt.Errorf("can't get messages: %w", err)
	}

	return c.JSON(messages)
}

//	@Summary		Обновить состояние сообщений
//	@Description	Обновляет состояние сообщений. Состояние обновляется индивидуально для каждого сообщения, игнорируя ошибки
//	@Security		MobileToken
//	@Tags			Устройство, Сообщения
//	@Accept			json
//	@Produce		json
//	@Param			request	body		[]smsgateway.MessageState	true	"Состояние сообщений"
//	@Success		204		{object}	nil							"Обновление выполнено"
//	@Failure		400		{object}	smsgateway.ErrorResponse	"Некорректный запрос"
//	@Failure		500		{object}	smsgateway.ErrorResponse	"Внутренняя ошибка сервера"
//	@Router			/mobile/v1/message [patch]
//
// Обновить состояние сообщений
func (h *mobileHandler) patchMessage(device models.Device, c *fiber.Ctx) error {
	req := []smsgateway.MessageState{}
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := h.Validator.Var(req, "required,dive"); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	for _, v := range req {
		err := h.messagesSvc.UpdateState(device.ID, v)
		if err != nil && !errors.Is(err, repositories.ErrMessageNotFound) {
			h.Logger.Error("Can't update message status", zap.Error(err))
		}
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *mobileHandler) authorize(handler func(models.Device, *fiber.Ctx) error) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Locals("token").(string)

		device, err := h.authSvc.AuthorizeDevice(token)
		if err != nil {
			h.Logger.Error("Can't authorize device", zap.Error(err))
			return fiber.ErrUnauthorized
		}

		return handler(device, c)
	}
}

func (h *mobileHandler) Register(router fiber.Router) {
	router = router.Group("/mobile/v1")

	router.Post("/device", limiter.New(), apikey.New(apikey.Config{
		Authorizer: func(token string) bool {
			return h.authSvc.AuthorizeRegistration(token) == nil
		},
	}), h.postDevice)

	router.Use(apikey.New(apikey.Config{
		Authorizer: func(token string) bool {
			return len(token) > 0
		},
	}))

	router.Patch("/device", h.authorize(h.patchDevice))

	router.Get("/message", h.authorize(h.getMessage))
	router.Patch("/message", h.authorize(h.patchMessage))
}

type MobileHandlerParams struct {
	fx.In

	Logger    *zap.Logger
	Validator *validator.Validate

	AuthSvc     *auth.Service
	MessagesSvc *services.MessagesService
}

func newMobileHandler(params MobileHandlerParams) *mobileHandler {
	idGen, _ := nanoid.Standard(21)

	return &mobileHandler{
		Handler:     Handler{Logger: params.Logger, Validator: params.Validator},
		authSvc:     params.AuthSvc,
		messagesSvc: params.MessagesSvc,
		idGen:       idGen,
	}
}
