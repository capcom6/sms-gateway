package smsgateway

// Device registration response
type MobileRegisterResponse struct {
	Id       string `json:"id" example:"QslD_GefqiYV6RQXdkM6V"`    // New device ID
	Token    string `json:"token" example:"bP0ZdK6rC6hCYZSjzmqhQ"` // Device access token
	Login    string `json:"login" example:"VQ4GII"`                // User login
	Password string `json:"password" example:"cp2pydvxd2zwpx"`     // User password
}

// Error response
type ErrorResponse struct {
	Message string `json:"message" example:"An error occurred"` // Error message
	Code    int32  `json:"code,omitempty"`                      // Error code
	Data    any    `json:"data,omitempty"`                      // Error context
}
