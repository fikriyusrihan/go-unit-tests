package controllers

type AppController interface {
	AuthController
}

type appController struct {
	AuthController
}

func NewAppController(
	authController AuthController,
) AppController {
	return &appController{
		AuthController: authController,
	}
}
