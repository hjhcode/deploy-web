package authv1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hjhcode/deploy-web/managers"
	"github.com/hjhcode/deploy-web/router/controllers/base"
)

func RegisterHost(router *gin.RouterGroup) {
	router.POST("host/add", httpHandlerHostAdd)
	router.POST("host/del", httpHandlerHostDel)
	router.POST("host/update", httpHandlerHostUpdate)
	router.GET("host/show", httpHandlerHostShow)
	router.GET("host/search", httpHandlerHostSearch)
}

type HostParam struct {
}

func httpHandlerHostAdd(c *gin.Context) {

}

func httpHandlerHostDel(c *gin.Context) {

}

func httpHandlerHostUpdate(c *gin.Context) {

}

func httpHandlerHostShow(c *gin.Context) {
	hostList, num := managers.GetAllHost()
	response := map[string]interface{}{
		"total_page": num,
		"datas":      hostList,
	}
	c.JSON(http.StatusOK, base.Success(response))
}

func httpHandlerHostSearch(c *gin.Context) {

}
