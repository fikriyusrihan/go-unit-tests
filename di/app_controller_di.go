package di

import (
	"go-product/controllers"
	"go-product/infrastructure/db"
	"go-product/repositories"
)

func InitializeAppController() (controllers.AppController, error) {
	database, err := db.NewPostgresDB()
	if err != nil {
		return nil, err
	}

	userRepository := repositories.NewUserRepository(database)
	productRepository := repositories.NewProductRepository(database)

	authController := controllers.NewAuthController(userRepository)
	productController := controllers.NewProductController(userRepository, productRepository)

	return controllers.NewAppController(authController, productController), nil
}
