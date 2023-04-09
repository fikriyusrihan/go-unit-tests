package di

import (
	"go-product/controllers"
	"go-product/infrastructure/db"
	"go-product/repositories"
	"go-product/services"
)

func InitializeAppController() (controllers.AppController, error) {
	database, err := db.NewPostgresDB()
	if err != nil {
		return nil, err
	}

	userRepository := repositories.NewUserRepository(database)
	productRepository := repositories.NewProductRepository(database)

	authService := services.NewAuthService(userRepository)
	productService := services.NewProductService(productRepository)

	authController := controllers.NewAuthController(authService)
	productController := controllers.NewProductController(productService)

	return controllers.NewAppController(authController, productController), nil
}
