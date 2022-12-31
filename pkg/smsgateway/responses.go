package smsgateway

// Успешная регистрация устройства
type MobileRegisterResponse struct {
	Id       string `json:"id" example:"QslD_GefqiYV6RQXdkM6V"`    // Идентификатор
	Token    string `json:"token" example:"bP0ZdK6rC6hCYZSjzmqhQ"` // Ключ доступа
	Login    string `json:"login" example:"VQ4GII"`                // Логин пользователя
	Password string `json:"password" example:"cp2pydvxd2zwpx"`     // Пароль пользователя
}

// Сообщение об ошибке
type ErrorResponse struct {
	Message string `json:"message" example:"Произошла ошибка"` // текст ошибки
	Code    int32  `json:"code,omitempty"`                     // код ошибки
	Data    any    `json:"data,omitempty"`                     // контекст
}
