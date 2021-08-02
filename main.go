package main

import (
	"fmt"
	"log"
	"os"

	"mygra.tech/project1/Config"
	"mygra.tech/project1/Models"
	"mygra.tech/project1/Routes"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var err error

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	// Creating a connection to the database
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))

	if err != nil {
		fmt.Println("Statuses: ", err)
	}

	defer Config.DB.Close()

	// Run the migrations: todo struc
	Config.DB.AutoMigrate(&Models.Todo{})

	// Setup routes
	r := Routes.SetupRouter()

	serverPort := os.Getenv("SERVER_PORT");

	// Running
	r.Run(":" + serverPort)
}