package Controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"mygra.tech/project1/Config"
	"mygra.tech/project1/Models"
	"mygra.tech/project1/Services"
	"mygra.tech/project1/Utils/Constants"
	"mygra.tech/project1/Utils/Formatters"
	"mygra.tech/project1/Utils/Helpers/Log"
	Utils "mygra.tech/project1/Utils/Paginations"
	"mygra.tech/project1/Utils/Responses"
)

type userController struct {
	service Services.UserService
}

func InitUserController(service Services.UserService) *userController {
	return &userController{service}
}

// List all users
func (controller *userController) GetUsers(c *gin.Context) {
	var responses Responses.ResponseApi

	pagination := Utils.GeneratePaginationFromRequest(c)
	result, err := controller.service.GetUsers(&pagination)

	if err != nil {
		responses = Formatters.Format(err, Constants.ERROR_RC500, Constants.ERROR_RM500)
		c.JSON(http.StatusOK, responses)
		return
	}

	responses = Formatters.Format(result, Constants.SUCCESS_RC200, Constants.SUCCESS_RM200)
	c.JSON(http.StatusOK, responses)
}

// Create a User
func (controller *userController) CreateAUser(c *gin.Context) {
	client, err := Config.InitRedis()
	if err != nil {
		Log.ERROR(err.Error())
	}
	var responses Responses.ResponseApi

	var user Models.User
	c.BindJSON(&user)

	result, err := controller.service.CreateAUser(user)

	if err != nil {
		responses = Formatters.Format(err, Constants.ERROR_RC500, Constants.ERROR_RM500)
		c.JSON(http.StatusOK, responses)
		return
	}

	responses = Formatters.Format(result, Constants.SUCCESS_RC200, Constants.SUCCESS_RM200)
	json, err := json.Marshal(responses)
	if err != nil {
		Log.ERROR(err.Error())
	}
	const name = "UserCreate"
	err = client.Set(name, json, 0).Err()

	if err != nil {
		Log.ERROR(err.Error())
	}

	val, err := client.Get(name).Result()
	if err != nil {
		Log.ERROR(err.Error())
	}
	Log.DEBUG(val)
	c.JSON(http.StatusOK, responses)
}

// Get a particular User with id
func (controller *userController) GetAUser(c *gin.Context) {
	var responses Responses.ResponseApi

	id := c.Params.ByName("id")
	result, err := controller.service.GetAUser(id)

	if err != nil {
		responses := Formatters.Format(err, Constants.ERROR_RC404, Constants.ERROR_RM404)

		c.JSON(http.StatusOK, responses)
		return
	}

	responses = Formatters.Format(result, Constants.SUCCESS_RC200, Constants.SUCCESS_RM200)
	c.JSON(http.StatusOK, responses)
}

// Update an existing User
func (controller *userController) UpdateAUser(c *gin.Context) {
	var responses Responses.ResponseApi

	var user Models.User
	e := c.BindJSON(&user)
	if e != nil {
		fmt.Println(e)
	}

	id := c.Params.ByName("id")
	result, err := controller.service.UpdateAUser(user, id)
	if err != nil {
		responses = Formatters.Format(err, Constants.ERROR_RC500, Constants.ERROR_RM500)
		c.JSON(http.StatusOK, responses)
		return
	}

	responses = Formatters.Format(result, Constants.SUCCESS_RC200, Constants.SUCCESS_RM200)
	c.JSON(http.StatusOK, responses)
}

// Delete a User
func (controller *userController) DeleteAUser(c *gin.Context) {
	var responses Responses.ResponseApi
	var user Models.User
	id := c.Params.ByName("id")

	user, err := controller.service.GetAUser(id)
	if err != nil {
		responses = Formatters.Format(err, Constants.ERROR_RC404, Constants.ERROR_RM404)
		c.JSON(http.StatusOK, responses)
		return
	}

	errDelete := controller.service.DeleteAUser(user, id)

	if errDelete != nil {
		responses = Formatters.Format(errDelete, Constants.ERROR_RC404, Constants.ERROR_RM404)
		c.JSON(http.StatusOK, responses)
		return
	}

	responses = Formatters.Format(id, Constants.SUCCESS_RC200, Constants.SUCCESS_RM200)
	c.JSON(http.StatusOK, responses)
}
