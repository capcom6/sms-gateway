package smsgateway

type MobileRegisterResponse struct {
	Id       string `json:"id"`
	Token    string `json:"token"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
