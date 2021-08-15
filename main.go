package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"mygra.tech/project1/Config"
	"mygra.tech/project1/Routes"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	Config.InitCassandra()
	
	db := Config.DatabaseOpen()

	// Setup routes
	r := Routes.SetupRouter(db)

	// Setup port
	serverPort := os.Getenv("SERVER_PORT")

	// Running
	r.Run(":" + serverPort)
}
