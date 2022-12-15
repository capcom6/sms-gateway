package smsgateway

// Запрос на регистрацию устройства
type MobileRegisterRequest struct {
	Name      *string `json:"name,omitempty" validate:"omitempty,max=128" example:"Android Phone"`    // Имя устройства
	PushToken *string `json:"pushToken" validate:"omitempty,max=256" example:"gHz-T6NezDlOfllr7F-Be"` // Токен для отправки PUSH-уведомлений
}
