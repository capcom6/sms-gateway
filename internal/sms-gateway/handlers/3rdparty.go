package handlers

import (
	"errors"
	"fmt"

	"github.com/capcom6/sms-gateway/internal/sms-gateway/models"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/modules/auth"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/modules/health"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/modules/messages"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/repositories"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/services"
	"github.com/capcom6/sms-gateway/pkg/maps"
	"github.com/capcom6/sms-gateway/pkg/smsgateway"
	"github.com/capcom6/sms-gateway/pkg/types"
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

	AuthSvc     *auth.Service
	MessagesSvc *messages.Service
	DevicesSvc  *services.DevicesService
	HealthSvc   *health.Service

	Logger    *zap.Logger
	Validator *validator.Validate
}

type thirdPartyHandler struct {
	Handler

	authSvc     *auth.Service
	messagesSvc *messages.Service
	devicesSvc  *services.DevicesService
	healthSvc   *health.Service
}

//	@Summary		Health check
//	@Description	Checks if service is healthy
//	@Tags			User
//	@Produce		json
//	@Success		200	{object}	smsgateway.HealthResponse	"Health check result"
//	@Failure		500	{object}	smsgateway.HealthResponse	"Service is unhealthy"
//	@Router			/3rdparty/v1/health [get]
//
// Health check
func (h *thirdPartyHandler) getHealth(c *fiber.Ctx) error {
	check, err := h.healthSvc.HealthCheck(c.Context())
	if err != nil {
		return err
	}

	res := smsgateway.HealthResponse{
		Status: smsgateway.HealthStatus(check.Status),
		Checks: maps.MapValues(
			check.Checks,
			func(c health.CheckDetail) smsgateway.HealthCheck {
				return smsgateway.HealthCheck{
					Description:   c.Description,
					ObservedUnit:  c.ObservedUnit,
					ObservedValue: c.ObservedValue,
					Status:        smsgateway.HealthStatus(c.Status),
				}
			},
		),
	}

	if check.Status == health.StatusFail {
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

//	@Summary		List devices
//	@Description	Returns list of registered devices
//	@Security		ApiAuth
//	@Tags			User
//	@Produce		json
//	@Success		200	{object}	[]smsgateway.Device			"Device list"
//	@Failure		400	{object}	smsgateway.ErrorResponse	"Invalid request"
//	@Failure		401	{object}	smsgateway.ErrorResponse	"Unauthorized"
//	@Failure		500	{object}	smsgateway.ErrorResponse	"Internal server error"
//	@Router			/3rdparty/v1/device [get]
//
// List devices
func (h *thirdPartyHandler) getDevice(user models.User, c *fiber.Ctx) error {
	devices, err := h.devicesSvc.Select(user)
	if err != nil {
		return fmt.Errorf("can't select devices: %w", err)
	}

	response := make([]smsgateway.Device, 0, len(devices))

	for _, device := range devices {
		response = append(response, smsgateway.Device{
			ID:        device.ID,
			Name:      types.OrDefault[string](device.Name, ""),
			CreatedAt: device.CreatedAt,
			UpdatedAt: device.UpdatedAt,
			DeletedAt: device.DeletedAt,
			LastSeen:  device.LastSeen,
		})
	}

	return c.JSON(response)
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
	if err := req.Validate(); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	skipPhoneValidation := c.QueryBool("skipPhoneValidation", false)

	devices, err := h.devicesSvc.Select(user)
	if err != nil {
		return fmt.Errorf("can't select devices: %w", err)
	}

	if len(devices) < 1 {
		return fiber.NewError(fiber.StatusBadRequest, "Нет ни одного устройства в учетной записи")
	}

	device := devices[0]
	state, err := h.messagesSvc.Enqeue(device, req, messages.EnqueueOptions{SkipPhoneValidation: skipPhoneValidation})
	if err != nil {
		var err400 messages.ErrValidation
		if errors.As(err, &err400) {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		return err
	}

	location, err := c.GetRouteURL(route3rdPartyGetMessage, fiber.Map{
		"id": state.ID,
	})
	if err != nil {
		h.Logger.Error("Failed to get route URL", zap.String("route", route3rdPartyGetMessage), zap.Error(err))
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
//	@Router			/3rdparty/v1/message [get]
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

	router.Get("/health", h.getHealth)

	router.Use(basicauth.New(basicauth.Config{
		Authorizer: func(username string, password string) bool {
			return len(username) > 0 && len(password) > 0
		},
	}))

	router.Get("/device", h.authorize(h.getDevice))

	router.Post("/message", h.authorize(h.postMessage))
	router.Get("/message/:id", h.authorize(h.getMessage)).Name(route3rdPartyGetMessage)
}

func newThirdPartyHandler(params ThirdPartyHandlerParams) *thirdPartyHandler {
	return &thirdPartyHandler{
		Handler:     Handler{Logger: params.Logger.Named("ThirdPartyHandler"), Validator: params.Validator},
		authSvc:     params.AuthSvc,
		messagesSvc: params.MessagesSvc,
		devicesSvc:  params.DevicesSvc,
		healthSvc:   params.HealthSvc,
	}
}
