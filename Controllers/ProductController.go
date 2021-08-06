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

type productController struct {
	service Services.ProductService
}

func InitProductController(service Services.ProductService) *productController {
	return &productController{service}
}

// List all products
func (controller *productController) GetProducts(c *gin.Context) {
	var responses Responses.ResponseApi

	pagination := Utils.GeneratePaginationFromRequest(c)
	result, err := controller.service.GetProducts(&pagination)

	if err != nil {
		responses = Formatters.Format(err, Constants.ERROR_RC500, Constants.ERROR_RM500)
		c.JSON(http.StatusOK, responses)
		return
	}

	responses = Formatters.Format(result, Constants.SUCCESS_RC200, Constants.SUCCESS_RM200)
	c.JSON(http.StatusOK, responses)
}

// Create a Product
func (controller *productController) CreateAProduct(c *gin.Context) {
	var responses Responses.ResponseApi

	var product Models.Product
	c.BindJSON(&product)

	result, err := controller.service.CreateAProduct(product)

	if err != nil {
		responses = Formatters.Format(err, Constants.ERROR_RC500, Constants.ERROR_RM500)
		c.JSON(http.StatusOK, responses)
		return
	}

	responses = Formatters.Format(result, Constants.SUCCESS_RC200, Constants.SUCCESS_RM200)
	c.JSON(http.StatusOK, responses)
}

// Get a particular Product with id
func (controller *productController) GetAProduct(c *gin.Context) {
	var responses Responses.ResponseApi

	id := c.Params.ByName("id")
	result, err := controller.service.GetAProduct(id)

	if err != nil {
		responses := Formatters.Format(err, Constants.ERROR_RC404, Constants.ERROR_RM404)

		c.JSON(http.StatusOK, responses)
		return
	}

	responses = Formatters.Format(result, Constants.SUCCESS_RC200, Constants.SUCCESS_RM200)
	c.JSON(http.StatusOK, responses)
}

// Update an existing Product
func (controller *productController) UpdateAProduct(c *gin.Context) {
	var responses Responses.ResponseApi

	var product Models.Product
	e := c.BindJSON(&product)
	if e != nil {
		fmt.Println(e)
	}

	id := c.Params.ByName("id")
	result, err := controller.service.UpdateAProduct(product, id)
	if err != nil {
		responses = Formatters.Format(err, Constants.ERROR_RC500, Constants.ERROR_RM500)
		c.JSON(http.StatusOK, responses)
		return
	}

	responses = Formatters.Format(result, Constants.SUCCESS_RC200, Constants.SUCCESS_RM200)
	c.JSON(http.StatusOK, responses)
}

// Delete a Product
func (controller *productController) DeleteAProduct(c *gin.Context) {
	var responses Responses.ResponseApi
	var product Models.Product
	id := c.Params.ByName("id")

	product, err := controller.service.GetAProduct(id)
	if err != nil {
		responses = Formatters.Format(err, Constants.ERROR_RC404, Constants.ERROR_RM404)
		c.JSON(http.StatusOK, responses)
		return
	}

	errDelete := controller.service.DeleteAProduct(product, id)

	if errDelete != nil {
		responses = Formatters.Format(errDelete, Constants.ERROR_RC404, Constants.ERROR_RM404)
		c.JSON(http.StatusOK, responses)
		return
	}

	responses = Formatters.Format(id, Constants.SUCCESS_RC200, Constants.SUCCESS_RM200)
	c.JSON(http.StatusOK, responses)
}
