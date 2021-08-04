package main

import (
	"os"

	"mygra.tech/project1/Routes"
)

var err error

func main() {
	// Setup routes
	r := Routes.SetupRouter()

	// Setup port
	serverPort := os.Getenv("SERVER_PORT");

	// Running
	r.Run(":" + serverPort)
}