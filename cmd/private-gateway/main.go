package main

import (
	privategateway "github.com/capcom6/sms-gateway/internal/private-gateway"
)

//	@title			Private SMS Gateway - API
//	@version		1.0.0
//	@description	Provide methods for working with private SMS gateway

//	@contact.name	Aleksandr Soloshenko
//	@contact.email	i@capcom.me

//	@securitydefinitions.apikey	MobileToken
//	@in							header
//	@name						Authorization
//	@description				Mobile app authorization token

//	@securitydefinitions.basic	ApiAuth
//	@description				External API authorization credentials

//	@host		localhost:3001
//	@schemes	http
//	@BasePath	/api
//
// Private SMS Gateway
func main() {
	privategateway.Run()
}
