package handlers

import (
	"errors"

	"bitbucket.org/capcom6/smsgatewaybackend/internal/smsgateway/models"
	"bitbucket.org/capcom6/smsgatewaybackend/internal/smsgateway/services"
	"bitbucket.org/capcom6/smsgatewaybackend/pkg/smsgateway"
	microbase "bitbucket.org/soft-c/gomicrobase"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

type thirdPartyHandler struct {
	microbase.Handler

	authSvc     *services.AuthService
	messagesSvc *services.MessagesService
}

// @Summary		Поставить сообщение в очередь
// @Description	Ставит сообщение в очередь на отправку. Если идентификатор не указан, то он будет сгенерирован автоматически
// @Security		ApiAuth
// @Tags			Пользователь, Сообщения
// @Accept			json
// @Produce		json
// @Param			request	body		smsgateway.Message			true	"Сообщение"
// @Success		201		{object}	smsgateway.MessageState		"Сообщение поставлено в очередь"
// @Failure		401		{object}	smsgateway.ErrorResponse	"Ошибка авторизации"
// @Failure		400		{object}	smsgateway.ErrorResponse	"Некорректный запрос"
// @Failure		500		{object}	smsgateway.ErrorResponse	"Внутренняя ошибка сервера"
// @Router			/3rdparty/v1/message [post]
func (h *thirdPartyHandler) postMessage(user models.User, c *fiber.Ctx) error {
	req := smsgateway.Message{}
	if err := h.BodyParserValidator(c, &req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if len(user.Devices) < 1 {
		return fiber.NewError(fiber.StatusBadRequest, "Нет ни одного устройтсва в учетной записи")
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

func (h *thirdPartyHandler) getMessage(user models.User, c *fiber.Ctx) error {
	// id := c.Params("id")

	return fiber.ErrNotImplemented
}

func (h *thirdPartyHandler) authorize(handler func(models.User, *fiber.Ctx) error) fiber.Handler {
	return func(c *fiber.Ctx) error {
		username := c.Locals("username").(string)
		password := c.Locals("password").(string)

		user, err := h.authSvc.AuthorizeUser(username, password)
		if err != nil {
			errorLog.Println(err)
			return fiber.ErrUnauthorized
		}

		return handler(user, c)
	}
}

func (h *thirdPartyHandler) register(router fiber.Router) {
	router.Use(basicauth.New(basicauth.Config{
		Authorizer: func(username string, password string) bool {
			return len(username) > 0 && len(password) > 0
		},
	}))

	router.Post("/message", h.authorize(h.postMessage))
	router.Get("/message/:id", h.authorize(h.getMessage))
}

func newThirdPartyHandler(validator *validator.Validate, authSvc *services.AuthService, messagesSvc *services.MessagesService) *thirdPartyHandler {
	return &thirdPartyHandler{
		Handler: microbase.Handler{
			Validator: validator,
		},
		authSvc:     authSvc,
		messagesSvc: messagesSvc,
	}
}
