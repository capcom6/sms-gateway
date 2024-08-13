package main

import (
	smsgateway "github.com/capcom6/sms-gateway/internal/sms-gateway"
)

//	@securitydefinitions.basic	ApiAuth

//	@securitydefinitions.apikey	MobileToken
//	@in							header
//	@name						Authorization
//	@description				Mobile device token

//	@title			SMS Gateway for Android™ API
//	@version		{APP_VERSION}
//	@description	This API provides programmatic access to sending SMS messages on Android devices. Features include sending SMS, checking message status, device management, webhook configuration, and system health checks.

//	@contact.name	Aleksandr Soloshenko
//	@contact.email	sms@capcom.me

//	@host		localhost:3000
//	@host		sms.capcom.me
//	@schemes	https
//	@BasePath	/api
//
// SMS Gateway for Android
func main() {
	smsgateway.Run()
}
