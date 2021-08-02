package main

import (
	"fmt"

	"mygra.tech/project1/Config"
	"mygra.tech/project1/Models"
	"mygra.tech/project1/Routes"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
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

	// Running
	r.Run()
}