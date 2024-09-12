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

//	@contact.name	SMSGate Support
//	@contact.email	support@sms-gate.app

//	@host		localhost:3000/api
//	@host		api.sms-gate.app
//	@schemes	https
//
// SMS Gateway for Android
func main() {
	smsgateway.Run()
}
