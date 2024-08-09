package main

import (
	"fmt"
	"log"
	"websocketchat/cmd/app"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load .env file")
	}
	mainApp := app.NewMainApp()
	log.Fatal(mainApp.App.Listen(":3000"))
}
