package Orders

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"mygra.tech/project1/Config"
	"mygra.tech/project1/Controllers"
	"mygra.tech/project1/Middlewares"
	"mygra.tech/project1/Repositories"
	"mygra.tech/project1/Services"
)

func TestOrderHttpMustHaveNonNegativeProductAmount(t *testing.T) {
	db := Config.DatabaseOpen()

	productRepository := Repositories.InitProductRepository(db)

	orderRepository := Repositories.InitOrderRepository(db)
	orderService := Services.InitOrderService(orderRepository, productRepository)
	orderController := Controllers.InitOrderController(orderService)

	var jsonStr = []byte(`{"userId": 3,"price": 100000,"status":"PENDING","productId": 1}`)

	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.POST("/orders", Middlewares.DBTransactionMiddleware(db), orderController.CreateAOrder)

	req, err := http.NewRequest(http.MethodPost, "/orders", bytes.NewBuffer(jsonStr))

	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}
	w := httptest.NewRecorder()

	go ProcessHttp(r, w, req, "Test")

	if w.Code == http.StatusOK {
		t.Log(w.Body)
		t.Logf("Expected to get status %d is same ast %d\n", http.StatusOK, w.Code)
	} else {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
}

func ProcessHttp(r *gin.Engine, w *httptest.ResponseRecorder, req *http.Request, s string) {
	for i := 0; i < 5; i++ {
		fmt.Println("Running ", s, ":", i)
		r.ServeHTTP(w, req)
	}
}
