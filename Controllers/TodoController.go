package Controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"mygra.tech/project1/Models"
	"mygra.tech/project1/Services"
)

type todoController struct {
	service Services.TodoService
}

func InitTodoController(service Services.TodoService) *todoController {
	return &todoController{service}
}

// List all todos
func (controller *todoController) GetTodos(c *gin.Context) {
	result, err := controller.service.GetTodos();

	if err != nil {
		fmt.Println("Error connectioin Database : ", err);
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, result)

}

// Create a Todo
func (controller *todoController) CreateATodo(c *gin.Context) {
	var todo Models.Todo
	c.BindJSON(&todo)

	result, err := controller.service.CreateATodo(todo)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, result)
}

// Get a particular Todo with id
func (controller *todoController) GetATodo(c *gin.Context) {
	id := c.Params.ByName("id")
	result, err := controller.service.GetATodo(id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, result)
}

// Update an existing Todo
func (controller *todoController) UpdateATodo(c *gin.Context) {
	var todo Models.Todo
	e := c.BindJSON(&todo)
	if e != nil {
		fmt.Println(e)
	}

	fmt.Println(todo)
	id := c.Params.ByName("id")
	result, err := controller.service.UpdateATodo(todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, result)
	}
}

// Delete a Todo
func (controller *todoController) DeleteATodo(c *gin.Context) {
	var todo Models.Todo
	id := c.Params.ByName("id")

	todo, err := controller.service.GetATodo(id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	
	err2 := controller.service.DeleteATodo(todo, id)

	if err2 != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
	}
}