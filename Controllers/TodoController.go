package Controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"mygra.tech/project1/Models"
	"mygra.tech/project1/Services"
	"mygra.tech/project1/Utils/Constants"
	"mygra.tech/project1/Utils/Formatters"
	"mygra.tech/project1/Utils/Responses"
)

type todoController struct {
	service Services.TodoService
}

func InitTodoController(service Services.TodoService) *todoController {
	return &todoController{service}
}

// List all todos
func (controller *todoController) GetTodos(c *gin.Context) {
	var responses Responses.ResponseApi
	result, err := controller.service.GetTodos();

	if err != nil {
		responses = Formatters.Format(err, Constants.ERROR_RC500, Constants.ERROR_RM500)
		c.JSON(http.StatusOK, responses)
		return
	}

	responses = Formatters.Format(result, Constants.SUCCESS_RC200, Constants.SUCCESS_RM200)
	c.JSON(http.StatusOK, responses)
}

// Create a Todo
func (controller *todoController) CreateATodo(c *gin.Context) {
	var responses Responses.ResponseApi

	var todo Models.Todo
	c.BindJSON(&todo)

	result, err := controller.service.CreateATodo(todo)

	if err != nil {
		responses = Formatters.Format(err, Constants.ERROR_RC500, Constants.ERROR_RM500)
		c.JSON(http.StatusOK, responses)
		return
	}

	responses = Formatters.Format(result, Constants.SUCCESS_RC200, Constants.SUCCESS_RM200)
	c.JSON(http.StatusOK, responses)
}

// Get a particular Todo with id
func (controller *todoController) GetATodo(c *gin.Context) {
	var responses Responses.ResponseApi

	id := c.Params.ByName("id")
	result, err := controller.service.GetATodo(id)

	if err != nil {
		responses := Formatters.Format(err, Constants.ERROR_RC404, Constants.ERROR_RM404)

		c.JSON(http.StatusOK, responses)
		return
	}

	responses = Formatters.Format(result, Constants.SUCCESS_RC200, Constants.SUCCESS_RM200)
	c.JSON(http.StatusOK, responses)
}

// Update an existing Todo
func (controller *todoController) UpdateATodo(c *gin.Context) {
	var responses Responses.ResponseApi
	
	var todo Models.Todo
	e := c.BindJSON(&todo)
	if e != nil {
		fmt.Println(e)
	}

	id := c.Params.ByName("id")
	result, err := controller.service.UpdateATodo(todo, id)
	if err != nil {
		responses = Formatters.Format(err, Constants.ERROR_RC500, Constants.ERROR_RM500)
		c.JSON(http.StatusOK, responses)
		return
	}

	responses = Formatters.Format(result, Constants.SUCCESS_RC200, Constants.SUCCESS_RM200)
	c.JSON(http.StatusOK, responses)	
}

// Delete a Todo
func (controller *todoController) DeleteATodo(c *gin.Context) {
	var responses Responses.ResponseApi
	var todo Models.Todo
	id := c.Params.ByName("id")

	todo, err := controller.service.GetATodo(id)
	if err != nil {
		responses = Formatters.Format(err, Constants.ERROR_RC404, Constants.ERROR_RM404)
		c.JSON(http.StatusOK, responses)
		return
	}
	
	errDelete := controller.service.DeleteATodo(todo, id)

	if errDelete != nil {
		responses = Formatters.Format(errDelete, Constants.ERROR_RC404, Constants.ERROR_RM404)
		c.JSON(http.StatusOK, responses)
		return
	}

	responses = Formatters.Format(id, Constants.SUCCESS_RC200, Constants.SUCCESS_RM200)
	c.JSON(http.StatusOK, responses)
}