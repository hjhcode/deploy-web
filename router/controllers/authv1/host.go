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
}

type HostParam struct {
	HostId   int64  `json:"host_id" form:"host_id"`
	HostName string `json:"host_name" form:"host_name" binding:"required"`
	HostIp   string `json:"host_ip" form:"host_ip" binding:"required"`
}

type HostIdParam struct {
	HostId int64 `json:"host_id" form:"host_id" binding:"required"`
}

func httpHandlerHostAdd(c *gin.Context) {
	var host HostParam
	err := c.Bind(&host)
	if err != nil {
		panic(err.Error())
	}

	result, mess := managers.AddNewHost(host.HostName, host.HostIp)
	if result {
		c.JSON(http.StatusOK, base.Success())
	} else {
		c.JSON(http.StatusOK, base.Fail(mess))
	}
}

func httpHandlerHostDel(c *gin.Context) {
	var host HostIdParam
	err := c.Bind(&host)
	if err != nil {
		panic(err.Error())
	}

	result := managers.DelHost(host.HostId)
	if result {
		c.JSON(http.StatusOK, base.Success())
	} else {
		c.JSON(http.StatusOK, base.Fail())
	}

}

func httpHandlerHostUpdate(c *gin.Context) {
	var host HostParam
	err := c.Bind(&host)
	if err != nil {
		panic(err.Error())
	}

	result := managers.UpdateHost(host.HostId, host.HostName, host.HostIp)
	if result {
		c.JSON(http.StatusOK, base.Success())
	} else {
		c.JSON(http.StatusOK, base.Fail())
	}
}

func httpHandlerHostShow(c *gin.Context) {
	hostList, num := managers.GetAllHost()
	response := map[string]interface{}{
		"total_page": num,
		"datas":      hostList,
	}
	c.JSON(http.StatusOK, base.Success(response))
}
