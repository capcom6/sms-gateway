package smsgateway

type MobileRegisterRequest struct {
	Name      string `json:"name,omitempty" validate:"omitempty,max=128"`
	PushToken string `json:"pushToken" validate:"omitempty,max=256"`
}
