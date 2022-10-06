package handlers

import (
	"errors"
	"fmt"
	"strings"

	"bitbucket.org/capcom6/smsgatewaybackend/internal/smsgateway/models"
	"bitbucket.org/capcom6/smsgatewaybackend/internal/smsgateway/repositories"
	"bitbucket.org/capcom6/smsgatewaybackend/internal/smsgateway/services"
	"bitbucket.org/capcom6/smsgatewaybackend/pkg/smsgateway"
	"bitbucket.org/soft-c/gohelpers/pkg/fiber/middleware/apikey"
	microbase "bitbucket.org/soft-c/gomicrobase"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jaevor/go-nanoid"
)

type mobileHandler struct {
	microbase.Handler

	authSvc     *services.AuthService
	messagesSvc *services.MessagesService

	idGen func() string
}

func (h *mobileHandler) postRegister(c *fiber.Ctx) error {
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

func (h *mobileHandler) getMessage(device models.Device, c *fiber.Ctx) error {
	messages, err := h.messagesSvc.SelectPending(device.ID)
	if err != nil {
		return fmt.Errorf("can't get messages: %w", err)
	}

	return c.JSON(messages)
}

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
			errorLog.Printf("Can't update message status: %s\n", err.Error())
		}
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *mobileHandler) authorize(handler func(models.Device, *fiber.Ctx) error) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Locals("token").(string)

		device, err := h.authSvc.AuthorizeDevice(token)
		if err != nil {
			errorLog.Println(err)
			return fiber.ErrUnauthorized
		}

		return handler(device, c)
	}
}

func (h *mobileHandler) register(router fiber.Router) {
	router.Post("/register", h.postRegister)

	router.Use(apikey.New(apikey.Config{
		Authorizer: func(token string) bool {
			return len(token) > 0
		},
	}))

	router.Get("/message", h.authorize(h.getMessage))
	router.Patch("/message", h.authorize(h.patchMessage))
}

func newMobileHandler(validator *validator.Validate, authSvc *services.AuthService, messagesSvc *services.MessagesService) *mobileHandler {
	idGen, _ := nanoid.Standard(21)

	return &mobileHandler{
		Handler:     microbase.Handler{Validator: validator},
		authSvc:     authSvc,
		messagesSvc: messagesSvc,
		idGen:       idGen,
	}
}
