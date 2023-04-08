package main

import (
	"go-product/controllers"
	"go-product/infrastructure/db"
	"go-product/infrastructure/http/router"
	"go-product/repository"
	"log"
)

const (
	ServerPort = ":8000"
)

func main() {
	database, err := db.NewPostgresDB()
	if err != nil {
		log.Fatalln("an error occurred while connecting to database: ", err)
		return
	}

	userRepository := repository.NewUserRepository(database)

	authController := controllers.NewAuthController(userRepository)
	appController := controllers.NewAppController(authController)

	log.Println("server will run on port: ", ServerPort)
	err = router.NewRouter(appController).Run(ServerPort)
	if err != nil {
		log.Fatalln("an error occurred while running the app: ", err)
		return
	}
}
