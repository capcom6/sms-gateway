package handlers

// func New(params Params) Result {
// 	return Result{
// 		Handlers: []http.ApiHanlder{},
// 	}
// }

// func Register(router fiber.Router, db *gorm.DB) error {
// 	cfg := config.GetConfig()

// 	users := repositories.NewUsersRepository(db)
// 	devices := repositories.NewDevicesRepository(db)
// 	messages := repositories.NewMessagesRepository(db)

// 	validator := validator.New()
// 	authSvc := services.NewAuthService(users, devices)
// 	pushSvc := services.NewPushService(cfg.FCM.CredentialsJSON)
// 	messagesSvc := services.NewMessagesService(pushSvc, messages)

// 	newMobileHandler(validator, authSvc, messagesSvc).register(router.Group("/mobile/v1"))
// 	newThirdPartyHandler(validator, authSvc, messagesSvc).register(router.Group("/3rdparty/v1"))

// 	return nil
// }
