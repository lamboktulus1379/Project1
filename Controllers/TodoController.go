package Controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"mygra.tech/project1/Models"
	"mygra.tech/project1/Services"
	"mygra.tech/project1/Utils/Constants"
	"mygra.tech/project1/Utils/Response"
)

type todoController struct {
	service Services.TodoService
}

func InitTodoController(service Services.TodoService) *todoController {
	return &todoController{service}
}

// List all todos
func (controller *todoController) GetTodos(c *gin.Context) {
	response := Response.ResponseApi{}
	result, err := controller.service.GetTodos();

	if err != nil {
		response.StatusCode = Constants.ERROR_RC500
		response.StatusMessage = Constants.ERROR_RM500
		response.Data = err
		c.JSON(http.StatusOK, response)
		return
	}

	response.StatusCode = Constants.SUCCESS_RC200
	response.StatusMessage = Constants.SUCCESS_RM200
	response.Data = result

	c.JSON(http.StatusOK, response)
}

// Create a Todo
func (controller *todoController) CreateATodo(c *gin.Context) {
	response := Response.ResponseApi{}

	var todo Models.Todo
	c.BindJSON(&todo)

	result, err := controller.service.CreateATodo(todo)

	if err != nil {
		response.StatusCode = Constants.ERROR_RC500
		response.StatusMessage = Constants.ERROR_RM500
		response.Data = err
		c.JSON(http.StatusOK, response)
		return
	}

	response.StatusCode = Constants.SUCCESS_RC200
	response.StatusMessage = Constants.SUCCESS_RM200
	response.Data = result

	c.JSON(http.StatusOK, response)
}

// Get a particular Todo with id
func (controller *todoController) GetATodo(c *gin.Context) {
	response := Response.ResponseApi{}

	id := c.Params.ByName("id")
	result, err := controller.service.GetATodo(id)

	if err != nil {
		response.StatusCode = Constants.ERROR_RC404
		response.StatusMessage = Constants.ERROR_RM404
		response.Data = err
		c.JSON(http.StatusOK, response)
		return
	}

	response.StatusCode = Constants.SUCCESS_RC200
	response.StatusMessage = Constants.SUCCESS_RM200
	response.Data = result

	c.JSON(http.StatusOK, response)
}

// Update an existing Todo
func (controller *todoController) UpdateATodo(c *gin.Context) {
	response := Response.ResponseApi{}
	
	var todo Models.Todo
	e := c.BindJSON(&todo)
	if e != nil {
		fmt.Println(e)
	}

	id := c.Params.ByName("id")
	result, err := controller.service.UpdateATodo(todo, id)
	if err != nil {
		response.StatusCode = Constants.ERROR_RC500
		response.StatusMessage = Constants.ERROR_RM500
		response.Data = result
		c.JSON(http.StatusOK, response)
		return
	}

	response.StatusCode = Constants.SUCCESS_RC200
	response.StatusMessage = Constants.SUCCESS_RM200
	response.Data = result

	c.JSON(http.StatusOK, response)	
}

// Delete a Todo
func (controller *todoController) DeleteATodo(c *gin.Context) {
	response := Response.ResponseApi{}
	var todo Models.Todo
	id := c.Params.ByName("id")

	todo, err := controller.service.GetATodo(id)
	if err != nil {
		response.StatusCode = Constants.ERROR_RC404
		response.StatusMessage = Constants.ERROR_RM404
		response.Data = err
		c.JSON(http.StatusOK, response)
		return
	}
	
	errDelete := controller.service.DeleteATodo(todo, id)

	if errDelete != nil {
		response.StatusCode = Constants.ERROR_RC500
		response.StatusMessage = Constants.ERROR_RM500
		response.Data = err
		c.JSON(http.StatusOK, response)
		return
	}

	c.JSON(http.StatusOK, gin.H{"id: " + id: "deleted"})
}