package Controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
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
	tracer, closer := initJaeger("get-users")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	var responses Responses.ResponseApi

	pagination := Utils.GeneratePaginationFromRequest(c)
	result, err := controller.service.GetUsers(&pagination)

	if err != nil {
		responses = Formatters.Format(err, Constants.ERROR_RC500, Constants.ERROR_RM500)
		c.JSON(http.StatusOK, responses)
		return
	}

	responses = Formatters.Format(result, Constants.SUCCESS_RC200, Constants.SUCCESS_RM200)
	responseFmt := fmt.Sprint(responses)

	span := tracer.StartSpan("get-users")
	span.SetTag("request-api-get-users", "v1/users")

	defer span.Finish()

	ctx := opentracing.ContextWithSpan(context.Background(), span)

	c.JSON(http.StatusOK, responses)
	printResponse(ctx, responseFmt)
}

func printResponse(ctx context.Context, response string) {
	span, _ := opentracing.StartSpanFromContext(ctx, "printResponse")
	defer span.Finish()

	span.LogFields(
		log.String("event", "print-response"),
		log.String("value", response),
	)

	// println(response)
	span.LogKV("event", "println")
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

// initJaeger returns an instance of Jaeger Tracer that samples 100% of traces and logs all spans to stdout.
func initJaeger(service string) (opentracing.Tracer, io.Closer) {
	cfg := &config.Configuration{
		ServiceName: service,
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
		},
	}
	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))

	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	return tracer, closer
}
