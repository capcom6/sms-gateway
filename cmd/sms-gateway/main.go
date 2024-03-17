package main

import (
	smsgateway "github.com/capcom6/sms-gateway/internal/sms-gateway"
)

//	@title			SMS Gateway - API
//	@version		1.0.0
//	@description	Provides API for (Android) SMS Gateway

//	@contact.name	Aleksandr Soloshenko
//	@contact.email	i@capcom.me

//	@securitydefinitions.apikey	MobileToken
//	@in							header
//	@name						Authorization
//	@description				Mobile device token

//	@securitydefinitions.basic	ApiAuth
//	@in							header
//	@description				End-user authentication key

//	@host		localhost:3000
//	@host		sms.capcom.me
//	@schemes	https
//	@BasePath	/api
//
// SMS Gateway
func main() {
	smsgateway.Run()
}
