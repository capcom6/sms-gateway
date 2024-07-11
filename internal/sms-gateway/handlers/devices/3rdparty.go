package devices

import (
	"fmt"

	"github.com/android-sms-gateway/client-go/smsgateway"
	"github.com/capcom6/go-helpers/slices"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/handlers/base"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/models"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/modules/auth"
	"github.com/capcom6/sms-gateway/internal/sms-gateway/modules/devices"
	"github.com/capcom6/sms-gateway/pkg/types"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type thirdPartyControllerParams struct {
	fx.In

	DevicesSvc *devices.Service

	Logger *zap.Logger
}

type ThirdPartyController struct {
	base.Handler

	devicesSvc *devices.Service
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
//	@Router			/3rdparty/v1/devices [get]
//
// List devices
func (h *ThirdPartyController) getDevices(user models.User, c *fiber.Ctx) error {
	devices, err := h.devicesSvc.Select(devices.WithUserID(user.ID))
	if err != nil {
		return fmt.Errorf("can't select devices: %w", err)
	}

	response := slices.Map(devices, func(device models.Device) smsgateway.Device {
		return smsgateway.Device{
			ID:        device.ID,
			Name:      types.OrDefault(device.Name, ""),
			CreatedAt: device.CreatedAt,
			UpdatedAt: device.UpdatedAt,
			DeletedAt: device.DeletedAt,
			LastSeen:  device.LastSeen,
		}
	})

	return c.JSON(response)
}

func (h *ThirdPartyController) Register(router fiber.Router) {
	router.Get("", auth.WithUser(h.getDevices))
}

func NewThirdPartyController(params thirdPartyControllerParams) *ThirdPartyController {
	return &ThirdPartyController{
		Handler: base.Handler{
			Logger: params.Logger.Named("devices"),
		},
		devicesSvc: params.DevicesSvc,
	}
}
