package smsgateway

import "time"

type ProcessState string

const (
	MessageStatePending   ProcessState = "Pending"   // В ожидании
	MessageStateProcessed ProcessState = "Processed" // Обработано
	MessageStateSent      ProcessState = "Sent"      // Отправлено
	MessageStateDelivered ProcessState = "Delivered" // Доставлено
	MessageStateFailed    ProcessState = "Failed"    // Ошибка
)

// Устройство
type Device struct {
	ID        string    `json:"id" example:"PyDmBQZZXYmyxMwED8Fzy"`                 // Идентификатор
	Name      string    `json:"name" example:"My Device"`                           // Название устройства
	CreatedAt time.Time `json:"createdAt" example:"2020-01-01T00:00:00Z"`           // Дата создания
	UpdatedAt time.Time `json:"updatedAt" example:"2020-01-01T00:00:00Z"`           // Дата обновления
	DeletedAt time.Time `json:"deletedAt,omitempty" example:"2020-01-01T00:00:00Z"` // Дата удаления

	LastSeen time.Time `json:"lastSeen" example:"2020-01-01T00:00:00Z"` // Последняя активность
}

// Сообщение
type Message struct {
	ID                 string   `json:"id,omitempty" validate:"omitempty,max=36" example:"PyDmBQZZXYmyxMwED8Fzy"`                          // Идентификатор
	Message            string   `json:"message" validate:"required,max=65535" example:"Hello World!"`                                      // Текст сообщения
	TTL                *uint64  `json:"ttl,omitempty" validate:"omitempty,min=5" example:"86400"`                                          // Время жизни сообщения в секундах
	SimNumber          *uint8   `json:"simNumber,omitempty" validate:"omitempty,max=3" example:"1"`                                        // Номер сим-карты
	WithDeliveryReport *bool    `json:"withDeliveryReport,omitempty" example:"true"`                                                       // Запрашивать отчет о доставке
	IsEncrypted        bool     `json:"isEncrypted,omitempty" example:"true"`                                                              // Зашифровано
	PhoneNumbers       []string `json:"phoneNumbers" validate:"required,min=1,max=100,dive,required,min=10,max=128" example:"79990001234"` // Получатели
}

// Состояние сообщения
type MessageState struct {
	ID          string           `json:"id,omitempty" validate:"omitempty,max=36" example:"PyDmBQZZXYmyxMwED8Fzy"` // Идентификатор
	State       ProcessState     `json:"state" validate:"required" example:"Pending"`                              // Состояние
	IsHashed    bool             `json:"isHashed" example:"false"`                                                 // Хэшировано
	IsEncrypted bool             `json:"isEncrypted" example:"false"`                                              // Зашифровано
	Recipients  []RecipientState `json:"recipients" validate:"required,min=1,dive"`                                // Детализация состояния по получателям
}

// Детализация состояния
type RecipientState struct {
	PhoneNumber string       `json:"phoneNumber" validate:"required,min=10,max=128" example:"79990001234"` // Номер телефона или первые 16 символов SHA256
	State       ProcessState `json:"state" validate:"required" example:"Pending"`                          // Состояние
	Error       *string      `json:"error,omitempty" example:"timeout"`                                    // Ошибка
}
