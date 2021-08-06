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
	Utils "mygra.tech/project1/Utils/Paginations"
	"mygra.tech/project1/Utils/Responses"
)

type orderController struct {
	service        Services.OrderService
	productService Services.ProductService
}

func InitOrderController(service Services.OrderService, productService Services.ProductService) *orderController {
	return &orderController{service, productService}
}

// List all orders
func (controller *orderController) GetOrders(c *gin.Context) {
	var responses Responses.ResponseApi

	pagination := Utils.GeneratePaginationFromRequest(c)
	result, err := controller.service.GetOrders(&pagination)

	if err != nil {
		responses = Formatters.Format(err, Constants.ERROR_RC500, Constants.ERROR_RM500)
		c.JSON(http.StatusOK, responses)
		return
	}

	responses = Formatters.Format(result, Constants.SUCCESS_RC200, Constants.SUCCESS_RM200)
	c.JSON(http.StatusOK, responses)
}

// Create a Order
func (controller *orderController) CreateAOrder(c *gin.Context) {
	var responses Responses.ResponseApi

	var order Models.Order
	c.BindJSON(&order)

	result, err := controller.service.CreateAOrder(order)

	if err != nil {
		responses = Formatters.Format(err, Constants.ERROR_RC500, Constants.ERROR_RM500)
		c.JSON(http.StatusOK, responses)
		return
	}

	responses = Formatters.Format(result, Constants.SUCCESS_RC200, Constants.SUCCESS_RM200)
	c.JSON(http.StatusOK, responses)
}

// Get a particular Order with id
func (controller *orderController) GetAOrder(c *gin.Context) {
	var responses Responses.ResponseApi

	id := c.Params.ByName("id")
	result, err := controller.service.GetAOrder(id)

	if err != nil {
		responses := Formatters.Format(err, Constants.ERROR_RC404, Constants.ERROR_RM404)

		c.JSON(http.StatusOK, responses)
		return
	}

	responses = Formatters.Format(result, Constants.SUCCESS_RC200, Constants.SUCCESS_RM200)
	c.JSON(http.StatusOK, responses)
}

// Update an existing Order
func (controller *orderController) UpdateAOrder(c *gin.Context) {
	var responses Responses.ResponseApi

	var order Models.Order
	e := c.BindJSON(&order)
	if e != nil {
		fmt.Println(e)
	}

	id := c.Params.ByName("id")
	result, err := controller.service.UpdateAOrder(order, id)
	if err != nil {
		responses = Formatters.Format(err, Constants.ERROR_RC500, Constants.ERROR_RM500)
		c.JSON(http.StatusOK, responses)
		return
	}

	responses = Formatters.Format(result, Constants.SUCCESS_RC200, Constants.SUCCESS_RM200)
	c.JSON(http.StatusOK, responses)
}

// Delete a Order
func (controller *orderController) DeleteAOrder(c *gin.Context) {
	var responses Responses.ResponseApi
	var order Models.Order
	id := c.Params.ByName("id")

	order, err := controller.service.GetAOrder(id)
	if err != nil {
		responses = Formatters.Format(err, Constants.ERROR_RC404, Constants.ERROR_RM404)
		c.JSON(http.StatusOK, responses)
		return
	}

	errDelete := controller.service.DeleteAOrder(order, id)

	if errDelete != nil {
		responses = Formatters.Format(errDelete, Constants.ERROR_RC404, Constants.ERROR_RM404)
		c.JSON(http.StatusOK, responses)
		return
	}

	responses = Formatters.Format(id, Constants.SUCCESS_RC200, Constants.SUCCESS_RM200)
	c.JSON(http.StatusOK, responses)
}
