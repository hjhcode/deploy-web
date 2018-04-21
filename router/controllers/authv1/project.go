package authv1

import "github.com/gin-gonic/gin"

func RegisterProject(router *gin.RouterGroup) {

	router.POST("project", httpHandlerAddProject)
	router.DELETE("project", httpHandlerDelProject)
	router.PUT("project", httpHandlerUpdateProject)
	router.GET("project", httpHandlerConstruct)
}

type NewDeploy struct {
	AccountId       int64
	ProjectName     string
	ProjectDescribe string
	GitDockerPath   string
	IsDel           int
	ProjectMember   string
}

type UpdateDeploy struct {
	ProjectId       int64
	AccountId       int64
	ProjectName     string
	ProjectDescribe string
	GitDockerPath   string
	IsDel           int
	ProjectMember   string
}

type JudgeAuthority struct {
	ProjectId int64
	AccountId int64
}

func httpHandlerAddProject(c *gin.Context) {

}

func httpHandlerDelProject(c *gin.Context) {

}

func httpHandlerUpdateProject(c *gin.Context) {

}

func httpHandlerConstruct(c *gin.Context) {

}
