package main

import (
	smsgateway "github.com/capcom6/sms-gateway/internal/sms-gateway"
)

//	@securitydefinitions.basic	ApiAuth

//	@securitydefinitions.apikey	MobileToken
//	@in							header
//	@name						Authorization
//	@description				Mobile device token

//	@title			SMS Gateway - API
//	@version		{APP_VERSION}
//	@description	Provides an API for (Android) SMS Gateway

//	@contact.name	Aleksandr Soloshenko
//	@contact.email	i@capcom.me

//	@host		localhost:3000
//	@host		sms.capcom.me
//	@schemes	https
//	@BasePath	/api
//
// SMS Gateway
func main() {
	smsgateway.Run()
}
