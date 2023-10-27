package main

import (
	"bitbucket.org/capcom6/smsgatewaybackend/internal/smsgateway"
)

//	@title			SMS-шлюз - API сервера
//	@version		1.0.0
//	@description	Предоставляет методы для взаимодействия с SMS-шлюзом

//	@contact.name	Aleksandr Soloshenko
//	@contact.email	capcom@soft-c.ru

//	@securitydefinitions.apikey	MobileToken
//	@in							header
//	@name						Authorization
//	@description				Авторизацию устройства по токену

//	@securitydefinitions.basic	ApiAuth
//	@description				Авторизацию пользователя по логин-паролю

//	@host		localhost:3000
//	@schemes	http
//	@BasePath	/api
//
// SMS-шлюз
func main() {
	smsgateway.Run()
}
