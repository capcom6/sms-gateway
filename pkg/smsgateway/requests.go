package smsgateway

// Device registration request
type MobileRegisterRequest struct {
	Name      *string `json:"name,omitempty" validate:"omitempty,max=128" example:"Android Phone"`    // Device name
	PushToken *string `json:"pushToken" validate:"omitempty,max=256" example:"gHz-T6NezDlOfllr7F-Be"` // FCM token
}

// Device update request
type MobileUpdateRequest struct {
	Id        string `json:"id" example:"QslD_GefqiYV6RQXdkM6V"`                                     // ID
	PushToken string `json:"pushToken" validate:"omitempty,max=256" example:"gHz-T6NezDlOfllr7F-Be"` // FCM token
}

// Push request
type UpstreamPushRequest = []PushNotification
