package Middlewares

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func StatusInList(status int, statusList []int) bool {
	for _, i := range statusList {
		if i == status {
			return true
		}
	}

	return false
}

func DBTransactionMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		trxHandle := db.Begin()
		log.Print("beginning database transaction")

		defer func() {
			if r := recover(); r != nil {
				trxHandle.Rollback()
			}
		}()

		c.Set("db_trx", trxHandle)
		c.Next()

		if StatusInList(c.Writer.Status(), []int{http.StatusOK, http.StatusCreated}) {
			log.Print("committing transactions")
			if err := trxHandle.Commit().Error; err != nil {
				log.Print("trx commit error: ", err)
			}
		} else {
			log.Print("rolling back transaction due to status code: ", c.Writer)
			trxHandle.Rollback()
		}
	}
}
