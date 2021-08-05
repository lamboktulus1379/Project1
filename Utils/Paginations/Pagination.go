package Utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"mygra.tech/project1/Models"
)

func GeneratePaginationFromRequest(c *gin.Context) Models.Pagination {
	limit := 10
	page := 1
	sort := "created_at asc"
	query := c.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
			break
		case "page":
			page, _ = strconv.Atoi(queryValue)
			break
		case "sort":
			sort = queryValue
			break
		}
	}
	return Models.Pagination{
		Limit: limit,
		Page:  page,
		Sort:  sort,
	}

}
