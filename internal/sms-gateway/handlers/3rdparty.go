package handlers

import (
	"errors"
	"fmt"

	"github.com/android-sms-gateway/client-go/smsgateway"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/handlers/base"
	devicesCtrl "github.com/capcom6/sms-gateway/internal/sms-gateway/handlers/devices"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/handlers/logs"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/handlers/webhooks"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/models"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/modules/auth"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/modules/devices"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/modules/messages"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/repositories"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

const (
	route3rdPartyGetMessage = "3rdparty.get.message"
)

type ThirdPartyHandlerParams struct {
	fx.In

	HealthHandler   *healthHandler
	WebhooksHandler *webhooks.ThirdPartyController
	DevicesHandler  *devicesCtrl.ThirdPartyController
	LogsHandler     *logs.ThirdPartyController

	AuthSvc     *auth.Service
	MessagesSvc *messages.Service
	DevicesSvc  *devices.Service

	Logger    *zap.Logger
	Validator *validator.Validate
}

type thirdPartyHandler struct {
	base.Handler

	healthHandler   *healthHandler
	webhooksHandler *webhooks.ThirdPartyController
	devicesHandler  *devicesCtrl.ThirdPartyController
	logsHandler     *logs.ThirdPartyController

	authSvc     *auth.Service
	messagesSvc *messages.Service
	devicesSvc  *devices.Service
}

//	@Summary		Enqueue message
//	@Description	Enqueues message for sending. If ID is not specified, it will be generated
//	@Security		ApiAuth
//	@Tags			User, Messages
//	@Accept			json
//	@Produce		json
//	@Param			skipPhoneValidation	query		bool						false	"Skip phone validation"
//	@Param			request				body		smsgateway.Message			true	"Send message request"
//	@Success		202					{object}	smsgateway.MessageState		"Message enqueued"
//	@Failure		400					{object}	smsgateway.ErrorResponse	"Invalid request"
//	@Failure		401					{object}	smsgateway.ErrorResponse	"Unauthorized"
//	@Failure		409					{object}	smsgateway.ErrorResponse	"Message with such ID already exists"
//	@Failure		500					{object}	smsgateway.ErrorResponse	"Internal server error"
//	@Header			202					{string}	Location					"Get message state URL"
//	@Router			/3rdparty/v1/message [post]
//
// Enqueue message
func (h *thirdPartyHandler) postMessage(user models.User, c *fiber.Ctx) error {
	req := smsgateway.Message{}
	if err := h.BodyParserValidator(c, &req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	skipPhoneValidation := c.QueryBool("skipPhoneValidation", false)

	devices, err := h.devicesSvc.Select(devices.WithUserID(user.ID))
	if err != nil {
		h.Logger.Error("Failed to select devices", zap.Error(err), zap.String("user_id", user.ID))
		return fiber.NewError(fiber.StatusInternalServerError, "Can't select devices. Please contact support")
	}

	if len(devices) < 1 {
		return fiber.NewError(fiber.StatusBadRequest, "No devices registered")
	}

	device := devices[0]
	state, err := h.messagesSvc.Enqeue(device, req, messages.EnqueueOptions{SkipPhoneValidation: skipPhoneValidation})
	if err != nil {
		var errValidation messages.ErrValidation
		if isBadRequest := errors.As(err, &errValidation); isBadRequest {
			return fiber.NewError(fiber.StatusBadRequest, errValidation.Error())
		}
		if isConflict := errors.Is(err, messages.ErrMessageAlreadyExists); isConflict {
			return fiber.NewError(fiber.StatusConflict, err.Error())
		}

		return fmt.Errorf("can't enqueue message: %w", err)
	}

	location, err := c.GetRouteURL(route3rdPartyGetMessage, fiber.Map{
		"id": state.ID,
	})
	if err != nil {
		h.Logger.Warn("Failed to get route URL", zap.String("route", route3rdPartyGetMessage), zap.Error(err))
	} else {
		c.Location(location)
	}

	return c.Status(fiber.StatusAccepted).JSON(state)
}

//	@Summary		Get message state
//	@Description	Returns message state by ID
//	@Security		ApiAuth
//	@Tags			User, Messages
//	@Produce		json
//	@Param			id	path		string						true	"Message ID"
//	@Success		200	{object}	smsgateway.MessageState		"Message state"
//	@Failure		400	{object}	smsgateway.ErrorResponse	"Invalid request"
//	@Failure		401	{object}	smsgateway.ErrorResponse	"Unauthorized"
//	@Failure		500	{object}	smsgateway.ErrorResponse	"Internal server error"
//	@Router			/3rdparty/v1/message/{id} [get]
//
// Get message state
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

func (h *thirdPartyHandler) Register(router fiber.Router) {
	router = router.Group("/3rdparty/v1")

	h.healthHandler.Register(router)

	router.Use(basicauth.New(basicauth.Config{
		Authorizer: func(username string, password string) bool {
			return len(username) > 0 && len(password) > 0
		},
	}), func(c *fiber.Ctx) error {
		username := c.Locals("username").(string)
		password := c.Locals("password").(string)

		user, err := h.authSvc.AuthorizeUser(username, password)
		if err != nil {
			h.Logger.Error("failed to authorize user", zap.Error(err))
			return fiber.ErrUnauthorized
		}

		c.Locals("user", user)

		return c.Next()
	})

	router.Post("/message", auth.WithUser(h.postMessage))
	router.Get("/message/:id", auth.WithUser(h.getMessage)).Name(route3rdPartyGetMessage)

	h.devicesHandler.Register(router.Group("/device")) // TODO: remove after 2025-07-11
	h.devicesHandler.Register(router.Group("/devices"))
	h.webhooksHandler.Register(router.Group("/webhooks"))
	h.logsHandler.Register(router.Group("/logs"))
}

func newThirdPartyHandler(params ThirdPartyHandlerParams) *thirdPartyHandler {
	return &thirdPartyHandler{
		Handler:         base.Handler{Logger: params.Logger.Named("ThirdPartyHandler"), Validator: params.Validator},
		healthHandler:   params.HealthHandler,
		webhooksHandler: params.WebhooksHandler,
		devicesHandler:  params.DevicesHandler,
		logsHandler:     params.LogsHandler,
		authSvc:         params.AuthSvc,
		messagesSvc:     params.MessagesSvc,
		devicesSvc:      params.DevicesSvc,
	}
}
