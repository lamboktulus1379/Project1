package Controllers

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Typing struct {
	Id      uuid.UUID `json:"id"`
	Content string    `json:"content"`
}

func GetRandom(c *gin.Context) {

	var s string = generateRandomLowerCaseAlphabet()

	uuidWithHyphen := uuid.New()
	var typing = Typing{Id: uuidWithHyphen, Content: s}
	c.JSON(http.StatusOK, typing)
}

func generateRandomLowerCaseAlphabet() string {
	rand.Seed(time.Now().UnixNano())

	var s string

	min := 97
	max := 123
	for i := 0; i < max; i++ {
		if (i+1)%5 == 0 {
			s += " "
		}

		v := rand.Intn(max-min) + min
		s += string(rune(v))
	}
	return s
}
