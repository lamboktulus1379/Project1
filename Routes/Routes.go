package Routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"mygra.tech/project1/Controllers"
	"mygra.tech/project1/Repositories"
	"mygra.tech/project1/Services"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	todoRepository := Repositories.InitTodoRepository(db)
	todoService := Services.InitTodoService(todoRepository)
	todoController := Controllers.InitTodoController(todoService)
	TODO_PATH := "todos"

	v1 := r.Group("/v1")
	{
		v1.GET(TODO_PATH, todoController.GetTodos)
		v1.POST(TODO_PATH, todoController.CreateATodo)
		v1.GET(TODO_PATH + "/:id", todoController.GetATodo)
		v1.PUT(TODO_PATH + "/:id", todoController.UpdateATodo)
		v1.DELETE(TODO_PATH + "/:id", todoController.DeleteATodo)
	}

	return r
}