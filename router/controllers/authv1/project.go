package authv1

import "github.com/gin-gonic/gin"

func RegisterProject(router *gin.RouterGroup) {

	router.POST("project", httpHandlerAddProject)
	router.DELETE("project", httpHandlerDelProject)
	router.PUT("project", httpHandlerUpdateProject)
	router.GET("project", httpHandlerConstruct)
}

func httpHandlerAddProject(c *gin.Context) {

}

func httpHandlerDelProject(c *gin.Context) {

}

func httpHandlerUpdateProject(c *gin.Context) {

}

func httpHandlerConstruct(c *gin.Context) {

}
