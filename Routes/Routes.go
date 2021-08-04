package Routes

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"mygra.tech/project1/Config"
	"mygra.tech/project1/Controllers"
	"mygra.tech/project1/Models"
	"mygra.tech/project1/Repositories"
	"mygra.tech/project1/Services"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	// Creating a connection to the database
	db, err := gorm.Open(os.Getenv("DB_TYPE"), Config.DbURL(Config.BuildDBConfig()));
	if err != nil {
		fmt.Println("Statuses: ", err)
	}

	// defer db.Close()

	// Run the migrations: todo struct
	db.AutoMigrate(&Models.Todo{})

	todoRepository := Repositories.InitTodoRepository(db)
	todoService := Services.InitTodoService(todoRepository)
	todoController := Controllers.InitTodoController(todoService)

	v1 := r.Group("/v1")
	{
		v1.GET("todos", todoController.GetTodos)
		v1.POST("todos", todoController.CreateATodo)
		v1.GET("todos/:id", todoController.GetATodo)
		v1.PUT("todos/:id", todoController.UpdateATodo)
		v1.DELETE("todos/:id", todoController.DeleteATodo)
	}

	return r
}