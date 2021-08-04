package Config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"mygra.tech/project1/Models"
)

var DB *gorm.DB

type DBConfig struct {
	Host string
	Port int
	User string
	DBName string
	Password string
}

func DatabaseOpen() *gorm.DB {
	// Creating a connection to the database
	db, err := gorm.Open(os.Getenv("DB_TYPE"), DbURL(BuildDBConfig()));
	if err != nil {
		fmt.Println("Statuses: ", err)
	}

	// defer db.Close()

	// Run the migrations: todo struct
	db.AutoMigrate(&Models.Todo{})

	return db
}

func BuildDBConfig() *DBConfig {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	port, err := strconv.Atoi(os.Getenv("DB_PORT"))

	if err != nil {
		fmt.Println("error occurred", err)
	}

	dbConfig := DBConfig{
		Host: os.Getenv("DB_HOST"),
		Port: port,
		User: os.Getenv("DB_USER"),
		DBName: os.Getenv("DB_NAME"),
		Password: os.Getenv("DB_PASSWORD"),
	}

	return &dbConfig
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}