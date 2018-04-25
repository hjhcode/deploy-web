package authv1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hjhcode/deploy-web/managers"
	"github.com/hjhcode/deploy-web/router/controllers/base"
)

func RegisterDeploy(router *gin.RouterGroup) {
	router.GET("deploy/show", httpHandlerDeployShow)
}

func httpHandlerDeployShow(c *gin.Context) {
	nums := c.Query("size")
	size, _ := strconv.Atoi(nums)
	requestPage := c.Query("page")
	start, _ := strconv.Atoi(requestPage)
	deployList, num := managers.GetAllDeploy(size, start)
	if deployList == nil {
		c.JSON(http.StatusOK, base.Fail("No content at the moment"))
	} else {
		response := map[string]interface{}{
			"request_page": start,
			"total_page":   num,
			"datas":        deployList,
		}
		c.JSON(http.StatusOK, base.Success(response))
	}
}
