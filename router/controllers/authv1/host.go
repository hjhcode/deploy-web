package authv1

import "github.com/gin-gonic/gin"

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

}

func httpHandlerHostSearch(c *gin.Context) {

}
