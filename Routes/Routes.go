package Routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

	userRepository := Repositories.InitUserRepository(db)
	userService := Services.InitUserService(userRepository)
	userController := Controllers.InitUserController(userService)
	USER_PATH := "users"

	v1 := r.Group("/v1")
	{
		v1.GET(TODO_PATH, todoController.GetTodos)
		v1.POST(TODO_PATH, todoController.CreateATodo)
		v1.GET(TODO_PATH+"/:id", todoController.GetATodo)
		v1.PUT(TODO_PATH+"/:id", todoController.UpdateATodo)
		v1.DELETE(TODO_PATH+"/:id", todoController.DeleteATodo)

		v1.GET(USER_PATH, userController.GetUsers)
		v1.POST(USER_PATH, userController.CreateAUser)
		v1.GET(USER_PATH+"/:id", userController.GetAUser)
		v1.PUT(USER_PATH+"/:id", userController.UpdateAUser)
		v1.DELETE(USER_PATH+"/:id", userController.DeleteAUser)
	}

	return r
}
