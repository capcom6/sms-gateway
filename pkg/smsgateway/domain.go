package smsgateway

type ProcessState string

const (
	MessageStatePending   ProcessState = "Pending"
	MessageStateSent      ProcessState = "Sent"
	MessageStateDelivered ProcessState = "Delivered"
	MessageStateFailed    ProcessState = "Failed"
)

type Message struct {
	ID           string   `json:"id,omitempty" validate:"omitempty,max=36"`
	Message      string   `json:"message" validate:"required,max=256"`
	PhoneNumbers []string `json:"phoneNumbers" validate:"required,min=1,max=100,dive,required,min=10"`
}

type MessageState struct {
	ID         string           `json:"id,omitempty" validate:"omitempty,max=36"`
	State      ProcessState     `json:"state" validate:"required"`
	Recipients []RecipientState `json:"recipients" validate:"required,min=1,dive"`
}

type RecipientState struct {
	PhoneNumber string       `json:"phoneNumber" validate:"required,len=11"`
	State       ProcessState `json:"state" validate:"required"`
}
