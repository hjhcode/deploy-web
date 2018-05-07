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
	router.GET("construct/detail", httpHandlerConstructDetail)
	router.POST("construct/start", httpHandlerConstructStart)
}

type ConstructIdParam struct {
	ConstructId int64 `json:"construct_id" form:"construct_id" binding:"required"`
}

func httpHandlerConstructShow(c *gin.Context) {
	constructList, num := managers.GetAllConstructRecord()
	if constructList == nil {
		c.JSON(http.StatusOK, base.Fail("No content at the moment"))
	} else {
		response := map[string]interface{}{
			"total_page": num,
			"datas":      constructList,
		}
		c.JSON(http.StatusOK, base.Success(response))
	}
}

func httpHandlerConstructDetail(c *gin.Context) {
	id := c.Query("id")
	constructId, _ := strconv.ParseInt(id, 10, 64)
	list := managers.GetConstructProjectData(constructId)
	if list == nil {
		c.JSON(http.StatusOK, base.Fail())
		return
	}
	c.JSON(http.StatusOK, base.Success(list))
}

func httpHandlerConstructStart(c *gin.Context) {
	var construct ConstructIdParam
	err := c.Bind(&construct)
	if err != nil {
		panic(err.Error())
	}
	//accountId := base.UserId(c)
	var accountId int64 = 1
	result, mess := managers.StartConstructProject(accountId, construct.ConstructId)
	if !result {
		c.JSON(http.StatusOK, base.Fail(mess))
		return
	}

	c.JSON(http.StatusOK, base.Success())
}
