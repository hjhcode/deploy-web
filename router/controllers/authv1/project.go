package authv1

import "github.com/gin-gonic/gin"

func RegisterProject(router *gin.RouterGroup) {

	router.POST("project", httpHandlerAddProject)   //用户创建工程（项目所有成员）
	router.DELETE("project", httpHandlerDelProject) //用户删除工程（项目所有成员）
	router.PUT("project", httpHandlerUpdateProject) //用户编辑工程（项目所有成员）
	router.GET("project", httpHandlerConstruct)
}

func httpHandlerAddProject(c *gin.Context) {

}

//删除工程(只有创建者可以删除，将删除标志重置）
func httpHandlerDelProject(c *gin.Context) {
	userId, _ := c.Get("userId")
	//获取projectId
	//判断userId是否是创建者

}

func httpHandlerUpdateProject(c *gin.Context) {

}

func httpHandlerConstruct(c *gin.Context) {

}
