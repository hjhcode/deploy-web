package authv1

import (
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hjhcode/deploy-web/managers"
	"github.com/hjhcode/deploy-web/router/controllers/base"
)

func RegisterProject(router *gin.RouterGroup) {

	router.POST("project/add", httpHandlerProjectAdd)
	router.POST("project/del", httpHandlerProjectDel)
	router.POST("project/update", httpHandlerProjectUpdate)
	router.POST("project/construct", httpHandlerConstruct)
	router.GET("project/show", httpHandlerProjectShow)     //分页展示工程
	router.GET("project/search", httpHandlerProjectSearch) //根据工程名查询
	router.GET("project/detail", httpHandlerProjectDetail) //单个工程详情
}

type ProjectParam struct {
	ProjectId       int64  `form:"project_id" json:"project_id"`
	ProjectName     string `form:"project_name" json:"project_name" binding:"required"`
	ProjectDescribe string `form:"project_describe" json:"project_describe" binding:"required"`
	GitDockerPath   string `form:"git_docker_path" json:"git_docker_path" binding:"required"`
	ProjectMember   string `form:"project_member" json:"project_member" binding:"required"`
}

type ProjectIdParam struct {
	ProjectId int64 `json:"project_id" binding:"required"`
}

func httpHandlerProjectAdd(c *gin.Context) {
	var project ProjectParam
	err := c.BindJSON(&project)
	if err != nil {
		panic(err.Error())
	}
	accountId := base.UserId(c)
	//var accountId int64 = 2
	if result, mess := managers.AddNewProject(accountId, project.ProjectName, project.ProjectDescribe,
		project.GitDockerPath, project.ProjectMember); !result {
		c.JSON(http.StatusOK, base.Fail(mess))
		return
	}

	c.JSON(http.StatusOK, base.Success())
}

func httpHandlerProjectDel(c *gin.Context) {
	var project ProjectIdParam
	err := c.BindJSON(&project)
	if err != nil {
		panic(err.Error())
	}
	accountId := base.UserId(c)
	//var accountId int64 = 2
	if result, mess := managers.DelProject(project.ProjectId, accountId); !result {
		c.JSON(http.StatusOK, base.Fail(mess))
		return

	}

	c.JSON(http.StatusOK, base.Success())
}

func httpHandlerProjectUpdate(c *gin.Context) {
	var project ProjectParam
	err := c.BindJSON(&project)
	if err != nil {
		panic(err.Error())
	}
	accountId := base.UserId(c)
	//var accountId int64 = 2
	if result, mess := managers.UpdateProject(accountId, project.ProjectId, project.ProjectName,
		project.ProjectDescribe, project.GitDockerPath, project.ProjectMember); !result {
		c.JSON(http.StatusOK, base.Fail(mess))
		return
	}

	c.JSON(http.StatusOK, base.Success())
}

func httpHandlerConstruct(c *gin.Context) {
	var project ProjectIdParam
	err := c.BindJSON(&project)
	if err != nil {
		panic(err.Error())
	}
	accountId := base.UserId(c)
	//var accountId int64 = 2
	if result, mess := managers.ConstructProject(accountId, project.ProjectId); !result {
		c.JSON(http.StatusOK, base.Fail(mess))
		return
	}

	c.JSON(http.StatusOK, base.Success())
}

//url?size=ssss
func httpHandlerProjectShow(c *gin.Context) {
	nums := c.Query("size")
	//nums := c.Query("")
	size, _ := strconv.Atoi(nums)
	requestPage := c.Query("requestpage")
	//requestPage := c.Query("")
	start, _ := strconv.Atoi(requestPage)
	projectList, num := managers.GetAllProject(size, start)
	response := map[string]interface{}{
		"request_page": requestPage,
		"total_page":   num,
		"datas":        projectList,
	}

	c.JSON(http.StatusOK, base.Success(response))
}

func httpHandlerProjectSearch(c *gin.Context) {
	nums := c.Query("")
	size, _ := strconv.Atoi(nums)
	requestPage := c.Query("")
	start, _ := strconv.Atoi(requestPage)
	projectName := c.Query("")
	projectList, num := managers.GetProjectByParam(projectName, size, start)
	response := map[string]interface{}{
		"request_page": requestPage,
		"total_page":   num,
		"datas":        projectList,
	}

	c.JSON(http.StatusOK, base.Success(response))
}

func httpHandlerProjectDetail(c *gin.Context) {
	//id := c.Query("")
	id := c.Query("id")
	projectId, _ := strconv.ParseInt(id, 10, 64)
	project := managers.GetOneProject(projectId)
	c.JSON(http.StatusOK, base.Success(project))
}
