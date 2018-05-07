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
	router.GET("deploy/detail", httpHandlerDeployDetail)
	router.POST("deploy/start", httpHandlerDeployStart)
	router.POST("deploy/rollback", httpHandlerDeployBack)
}

type DeployIdParam struct {
	DeployId int64 `json:"deploy_id" form:"deploy_id" binding:"required"`
	GroupId  int   `json:"group_id"  form:"group_id"`
}

func httpHandlerDeployShow(c *gin.Context) {
	deployList, num := managers.GetAllDeploy()
	if deployList == nil {
		c.JSON(http.StatusOK, base.Fail("No content at the moment"))
	} else {
		response := map[string]interface{}{
			"total_page": num,
			"datas":      deployList,
		}
		c.JSON(http.StatusOK, base.Success(response))
	}
}

func httpHandlerDeployDetail(c *gin.Context) {
	id := c.Query("id")
	deployId, _ := strconv.ParseInt(id, 10, 64)
	list := managers.GetDeployServiceData(deployId)
	if list == nil {
		c.JSON(http.StatusOK, base.Fail())
		return
	}
	c.JSON(http.StatusOK, base.Success(list))
}

func httpHandlerDeployStart(c *gin.Context) {
	var deploy DeployIdParam
	err := c.Bind(&deploy)
	if err != nil {
		panic(err.Error())
	}
	//accountId := base.UserId(c)
	var accountId int64 = 1
	result, mess := managers.StartDeployService(accountId, deploy.DeployId, deploy.GroupId)
	if !result {
		c.JSON(http.StatusOK, base.Fail(mess))
		return
	}

	c.JSON(http.StatusOK, base.Success())
}

func httpHandlerDeployBack(c *gin.Context) {
	var deploy DeployIdParam
	err := c.Bind(&deploy)
	if err != nil {
		panic(err.Error())
	}
	//accountId := base.UserId(c)
	var accountId int64 = 1
	result, mess := managers.BackDeployService(accountId, deploy.DeployId)
	if !result {
		c.JSON(http.StatusOK, base.Fail(mess))
		return
	}

	c.JSON(http.StatusOK, base.Success())
}
