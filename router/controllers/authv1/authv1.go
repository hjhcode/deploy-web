package authv1

import (
	"github.com/gin-gonic/gin"
)

func Register(router *gin.RouterGroup) {
	RegisterAccount(router)
	RegisterService(router)
	RegisterProject(router)
	RegisterConstructRecord(router)
	RegisterDeploy(router)
	RegisterMirror(router)
	RegisterHost(router)
}
