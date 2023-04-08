package controllers

type AppController interface {
	AuthController
	ProductController
}

type appController struct {
	AuthController
	ProductController
}

func NewAppController(
	authController AuthController,
	productController ProductController,
) AppController {
	return &appController{
		AuthController:    authController,
		ProductController: productController,
	}
}
