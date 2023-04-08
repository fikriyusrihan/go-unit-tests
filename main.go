package main

import (
	"go-product/di"
	"go-product/infrastructure/http/router"
	"log"
)

const (
	ServerPort = ":8000"
)

func main() {
	appController, err := di.InitializeAppController()
	if err != nil {
		log.Fatalln("an error occurred while initializing the app: ", err)
		return
	}

	log.Println("server will run on port: ", ServerPort)
	err = router.NewRouter(appController).Run(ServerPort)
	if err != nil {
		log.Fatalln("an error occurred while running the app: ", err)
		return
	}
}
