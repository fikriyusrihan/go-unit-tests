package main

import (
	"go-product/config"
	"go-product/di"
	"go-product/infrastructure/http/router"
	"log"
)

func main() {
	config.ReadConfig()

	appController, err := di.InitializeAppController()
	if err != nil {
		log.Fatalln("An error occurred while initializing the app: ", err)
		return
	}

	ServerPort := ":" + config.C.Server.Port
	log.Println("server will run on port: ", ServerPort)
	err = router.NewRouter(appController).Run(ServerPort)
	if err != nil {
		log.Fatalln("An error occurred while running the app: ", err)
		return
	}
}
