package handlers

import (
	"errors"
	"fmt"
	"strings"

	"github.com/android-sms-gateway/client-go/smsgateway"
	"github.com/capcom6/go-infra-fx/http/apikey"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/handlers/base"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/handlers/webhooks"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/models"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/modules/auth"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/modules/messages"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/repositories"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/jaevor/go-nanoid"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type mobileHandler struct {
	base.Handler

	authSvc     *auth.Service
	messagesSvc *messages.Service

	webhooksCtrl *webhooks.MobileController

	idGen func() string
}

//	@Summary		Register device
//	@Description	Registers new device and returns credentials
//	@Tags			Device
//	@Accept			json
//	@Produce		json
//	@Param			request	body		smsgateway.MobileRegisterRequest	true	"Device registration request"
//	@Success		201		{object}	smsgateway.MobileRegisterResponse	"Device registered"
//	@Failure		400		{object}	smsgateway.ErrorResponse			"Invalid request"
//	@Failure		401		{object}	smsgateway.ErrorResponse			"Unauthorized (private mode only)"
//	@Failure		429		{object}	smsgateway.ErrorResponse			"Too many requests"
//	@Failure		500		{object}	smsgateway.ErrorResponse			"Internal server error"
//	@Router			/mobile/v1/device [post]
//
// Register device
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

	device, err := h.authSvc.RegisterDevice(user, req.Name, req.PushToken)
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

//	@Summary		Update device
//	@Description	Updates push token for device
//	@Security		MobileToken
//	@Tags			Device
//	@Accept			json
//	@Param			request	body	smsgateway.MobileUpdateRequest	true	"Device update request"
//	@Success		204		"Successfully updated"
//	@Failure		400		{object}	smsgateway.ErrorResponse	"Invalid request"
//	@Failure		403		{object}	smsgateway.ErrorResponse	"Forbidden (wrong device ID)"
//	@Failure		500		{object}	smsgateway.ErrorResponse	"Internal server error"
//	@Router			/mobile/v1/device [patch]
//
// Update device
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

//	@Summary		Get messages for sending
//	@Description	Returns list of pending messages
//	@Security		MobileToken
//	@Tags			Device, Messages
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		smsgateway.Message			"List of pending messages"
//	@Failure		500	{object}	smsgateway.ErrorResponse	"Internal server error"
//	@Router			/mobile/v1/message [get]
//
// Get messages for sending
func (h *mobileHandler) getMessage(device models.Device, c *fiber.Ctx) error {
	messages, err := h.messagesSvc.SelectPending(device.ID)
	if err != nil {
		return fmt.Errorf("can't get messages: %w", err)
	}

	return c.JSON(messages)
}

//	@Summary		Update message state
//	@Description	Updates message state
//	@Security		MobileToken
//	@Tags			Device, Messages
//	@Accept			json
//	@Produce		json
//	@Param			request	body		[]smsgateway.MessageState	true	"New message state"
//	@Success		204		{object}	nil							"Successfully updated"
//	@Failure		400		{object}	smsgateway.ErrorResponse	"Invalid request"
//	@Failure		500		{object}	smsgateway.ErrorResponse	"Internal server error"
//	@Router			/mobile/v1/message [patch]
//
// Update message state
func (h *mobileHandler) patchMessage(device models.Device, c *fiber.Ctx) error {
	req := []smsgateway.MessageState{}
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	for _, v := range req {
		if err := h.ValidateStruct(v); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		err := h.messagesSvc.UpdateState(device.ID, v)
		if err != nil && !errors.Is(err, repositories.ErrMessageNotFound) {
			h.Logger.Error("Can't update message status", zap.Error(err))
		}
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *mobileHandler) Register(router fiber.Router) {
	router = router.Group("/mobile/v1")

	router.Post("/device", limiter.New(), apikey.New(apikey.Config{
		Next: func(c *fiber.Ctx) bool { return h.authSvc.IsPublic() },
		Authorizer: func(token string) bool {
			return h.authSvc.AuthorizeRegistration(token) == nil
		},
	}), h.postDevice)

	router.Use(apikey.New(apikey.Config{
		Authorizer: func(token string) bool {
			return len(token) > 0
		},
	}), func(c *fiber.Ctx) error {
		token := c.Locals("token").(string)

		device, err := h.authSvc.AuthorizeDevice(token)
		if err != nil {
			h.Logger.Error("Can't authorize device", zap.Error(err))
			return fiber.ErrUnauthorized
		}

		c.Locals("device", device)

		return c.Next()
	})

	router.Patch("/device", auth.WithDevice(h.patchDevice))

	router.Get("/message", auth.WithDevice(h.getMessage))
	router.Patch("/message", auth.WithDevice(h.patchMessage))

	h.webhooksCtrl.Register(router.Group("/webhooks"))
}

type mobileHandlerParams struct {
	fx.In

	Logger    *zap.Logger
	Validator *validator.Validate

	AuthSvc     *auth.Service
	MessagesSvc *messages.Service

	WebhooksCtrl *webhooks.MobileController
}

func newMobileHandler(params mobileHandlerParams) *mobileHandler {
	idGen, _ := nanoid.Standard(21)

	return &mobileHandler{
		Handler:      base.Handler{Logger: params.Logger, Validator: params.Validator},
		authSvc:      params.AuthSvc,
		messagesSvc:  params.MessagesSvc,
		webhooksCtrl: params.WebhooksCtrl,
		idGen:        idGen,
	}
}
