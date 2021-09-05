package Routes

import (
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mygra.tech/project1/Controllers"
	"mygra.tech/project1/Middlewares"
	"mygra.tech/project1/Repositories"
	"mygra.tech/project1/Services"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/swaggo/gin-swagger/example/basic/docs"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.New()

	url := ginSwagger.URL(os.Getenv("SERVER_HOST") + ":" + os.Getenv("SERVER_PORT") + "/swagger/doc.json")

	todoRepository := Repositories.InitTodoRepository(db)
	todoService := Services.InitTodoService(todoRepository)
	todoController := Controllers.InitTodoController(todoService)
	TODO_PATH := "todos"

	userRepository := Repositories.InitUserRepository(db)
	userService := Services.InitUserService(userRepository)
	userController := Controllers.InitUserController(userService)
	USER_PATH := "users"

	productRepository := Repositories.InitProductRepository(db)
	productService := Services.InitProductService(productRepository)
	productController := Controllers.InitProductController(productService)
	PRODUCT_PATH := "products"

	orderRepository := Repositories.InitOrderRepository(db)
	orderService := Services.InitOrderService(orderRepository, productRepository)
	orderController := Controllers.InitOrderController(orderService)
	ORDER_PATH := "orders"

	RANDOM_PATH := "randoms"

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

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

		v1.GET(PRODUCT_PATH, productController.GetProducts)
		v1.POST(PRODUCT_PATH, productController.CreateAProduct)
		v1.GET(PRODUCT_PATH+"/:id", productController.GetAProduct)
		v1.PUT(PRODUCT_PATH+"/:id", productController.UpdateAProduct)
		v1.DELETE(PRODUCT_PATH+"/:id", productController.DeleteAProduct)

		v1.GET(ORDER_PATH, Middlewares.DBTransactionMiddleware(db), orderController.GetOrders)
		v1.POST(ORDER_PATH, Middlewares.DBTransactionMiddleware(db), orderController.CreateAOrder)
		v1.GET(ORDER_PATH+"/:id", orderController.GetAOrder)
		v1.PUT(ORDER_PATH+"/:id", orderController.UpdateAOrder)
		v1.DELETE(ORDER_PATH+"/:id", orderController.DeleteAOrder)

		v1.GET(RANDOM_PATH, Controllers.GetRandom)
	}

	return r
}
