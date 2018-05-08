package authv1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hjhcode/deploy-web/managers"
	"github.com/hjhcode/deploy-web/router/controllers/base"
)

func RegisterMirror(router *gin.RouterGroup) {
	router.GET("mirror/show", httpHandlerMirrorShow)
}

func httpHandlerMirrorShow(c *gin.Context) {
	mirrorList, num := managers.GetAllMirror()
	response := map[string]interface{}{
		"total_page": num,
		"datas":      mirrorList,
	}
	c.JSON(http.StatusOK, base.Success(response))
}
