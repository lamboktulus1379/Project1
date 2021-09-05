package Typings

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"mygra.tech/project1/Controllers"
)

func TestTypingGetRandomSentence(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.GET("/randoms", Controllers.GetRandom)

	req := httptest.NewRequest(http.MethodGet, "/randoms", nil)

	w := httptest.NewRecorder()

	go r.ServeHTTP(w, req)

	if w.Code == http.StatusOK {
		t.Log(w.Body)
		t.Log(req.Body)
		t.Logf("Exppected get status %d is same as %d\n", http.StatusOK, w.Code)
	} else {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}

}
