package Routes

import (
	"github.com/gin-gonic/gin"
	"mygra.tech/project1/Controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("todos", Controllers.GetTodos)
		v1.POST("todos", Controllers.CreateATodo)
		v1.GET("todos/:id", Controllers.GetATodo)
		v1.PUT("todos/:id", Controllers.UpdateATodo)
		v1.DELETE("todos/:id", Controllers.DeleteATodo)
	}

	return r
}