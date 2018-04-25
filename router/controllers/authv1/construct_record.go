package authv1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hjhcode/deploy-web/managers"
	"github.com/hjhcode/deploy-web/router/controllers/base"
)

func RegisterConstructRecord(router *gin.RouterGroup) {
	router.GET("construct/show", httpHandlerConstructShow)
}

func httpHandlerConstructShow(c *gin.Context) {
	nums := c.Query("size")
	size, _ := strconv.Atoi(nums)
	requestPage := c.Query("page")
	start, _ := strconv.Atoi(requestPage)
	constructList, num := managers.GetAllConstructRecord(size, start)
	if constructList == nil {
		c.JSON(http.StatusOK, base.Fail("No content at the moment"))
	} else {
		response := map[string]interface{}{
			"request_page": start,
			"total_page":   num,
			"datas":        constructList,
		}
		c.JSON(http.StatusOK, base.Success(response))
	}
}
