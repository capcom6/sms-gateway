package smsgateway

import (
	"fmt"
	"time"
)

const (
	ProcessingStatePending   ProcessingState = "Pending"   // Pending
	ProcessingStateProcessed ProcessingState = "Processed" // Processed (received by device)
	ProcessingStateSent      ProcessingState = "Sent"      // Sent
	ProcessingStateDelivered ProcessingState = "Delivered" // Delivered
	ProcessingStateFailed    ProcessingState = "Failed"    // Failed
)

var allProcessStates = map[ProcessingState]struct{}{
	ProcessingStatePending:   {},
	ProcessingStateProcessed: {},
	ProcessingStateSent:      {},
	ProcessingStateDelivered: {},
	ProcessingStateFailed:    {},
}

// Device
type Device struct {
	ID        string     `json:"id" example:"PyDmBQZZXYmyxMwED8Fzy"`                 // ID
	Name      string     `json:"name" example:"My Device"`                           // Name
	CreatedAt time.Time  `json:"createdAt" example:"2020-01-01T00:00:00Z"`           // Created at (read only)
	UpdatedAt time.Time  `json:"updatedAt" example:"2020-01-01T00:00:00Z"`           // Updated at (read only)
	DeletedAt *time.Time `json:"deletedAt,omitempty" example:"2020-01-01T00:00:00Z"` // Deleted at (read only)

	LastSeen time.Time `json:"lastSeen" example:"2020-01-01T00:00:00Z"` // Last seen at (read only)
}

// Message
type Message struct {
	ID                 string   `json:"id,omitempty" validate:"omitempty,max=36" example:"PyDmBQZZXYmyxMwED8Fzy"`                          // ID (if not set - will be generated)
	Message            string   `json:"message" validate:"required,max=65535" example:"Hello World!"`                                      // Content
	SimNumber          *uint8   `json:"simNumber,omitempty" validate:"omitempty,max=3" example:"1"`                                        // SIM card number (1-3)
	WithDeliveryReport *bool    `json:"withDeliveryReport,omitempty" example:"true"`                                                       // With delivery report
	IsEncrypted        bool     `json:"isEncrypted,omitempty" example:"true"`                                                              // Is encrypted
	PhoneNumbers       []string `json:"phoneNumbers" validate:"required,min=1,max=100,dive,required,min=10,max=128" example:"79990001234"` // Recipients (phone numbers)

	TTL        *uint64    `json:"ttl,omitempty" validate:"omitempty,min=5" example:"86400"` // Time to live in seconds (conflicts with `validUntil`)
	ValidUntil *time.Time `json:"validUntil,omitempty" example:"2020-01-01T00:00:00Z"`      // Valid until (conflicts with `ttl`)
}

func (m Message) Validate() error {
	if m.TTL != nil && m.ValidUntil != nil {
		return fmt.Errorf("%w: ttl and validUntil", ErrConflictFields)
	}

	return nil
}

// Message state
type MessageState struct {
	ID          string               `json:"id,omitempty" validate:"omitempty,max=36" example:"PyDmBQZZXYmyxMwED8Fzy"` // Message ID
	State       ProcessingState      `json:"state" validate:"required" example:"Pending"`                              // State
	IsHashed    bool                 `json:"isHashed" example:"false"`                                                 // Hashed
	IsEncrypted bool                 `json:"isEncrypted" example:"false"`                                              // Encrypted
	Recipients  []RecipientState     `json:"recipients" validate:"required,min=1,dive"`                                // Recipients states
	States      map[string]time.Time `json:"states"`                                                                   // History of states
}

func (m MessageState) Validate() error {
	for k := range m.States {
		if _, ok := allProcessStates[ProcessingState(k)]; !ok {
			return fmt.Errorf("invalid state value: %s", k)
		}
	}

	return nil
}

// Recipient state
type RecipientState struct {
	PhoneNumber string          `json:"phoneNumber" validate:"required,min=10,max=128" example:"79990001234"` // Phone number or first 16 symbols of SHA256 hash
	State       ProcessingState `json:"state" validate:"required" example:"Pending"`                          // State
	Error       *string         `json:"error,omitempty" example:"timeout"`                                    // Error (for `Failed` state)
}

// Push notification
type PushNotification struct {
	Token string `json:"token" validate:"required" example:"PyDmBQZZXYmyxMwED8Fzy"` // Device FCM token
}
