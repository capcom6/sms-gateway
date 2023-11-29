package smsgateway

type ProcessState string

const (
	MessageStatePending   ProcessState = "Pending"   // В ожидании
	MessageStateProcessed ProcessState = "Processed" // Обработано
	MessageStateSent      ProcessState = "Sent"      // Отправлено
	MessageStateDelivered ProcessState = "Delivered" // Доставлено
	MessageStateFailed    ProcessState = "Failed"    // Ошибка
)

// Сообщение
type Message struct {
	ID           string   `json:"id,omitempty" validate:"omitempty,max=36" example:"PyDmBQZZXYmyxMwED8Fzy"`                  // Идентификатор
	Message      string   `json:"message" validate:"required,max=256" example:"Hello World!"`                                // Текст сообщения
	TTL          *uint64  `json:"ttl,omitempty" validate:"omitempty,min=5" example:"86400"`                                  // Время жизни сообщения в секундах
	SimNumber    *uint8   `json:"simNumber,omitempty" validate:"omitempty,max=3" example:"1"`                                // Номер сим-карты
	PhoneNumbers []string `json:"phoneNumbers" validate:"required,min=1,max=100,dive,required,min=10" example:"79990001234"` // Получатели
}

// Состояние сообщения
type MessageState struct {
	ID         string           `json:"id,omitempty" validate:"omitempty,max=36" example:"PyDmBQZZXYmyxMwED8Fzy"` // Идентификатор
	State      ProcessState     `json:"state" validate:"required" example:"Pending"`                              // Состояние
	Recipients []RecipientState `json:"recipients" validate:"required,min=1,dive"`                                // Детализация состояния по получателям
}

// Детализация состояния
type RecipientState struct {
	PhoneNumber string       `json:"phoneNumber" validate:"required,min=10" example:"79990001234"` // Номер телефона
	State       ProcessState `json:"state" validate:"required" example:"Pending"`                  // Состояние
	Error       *string      `json:"error,omitempty" example:"timeout"`                            // Ошибка
}
