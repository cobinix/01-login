package main

import (
	"log"

	"github.com/cobinix/01-login/app"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Print("Error loading .env file")
	}

	app.Init()
	StartServer()

}
