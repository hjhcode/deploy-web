package middleware

import (
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hjhcode/deploy-web/common/components"
	"github.com/hjhcode/deploy-web/router/controllers/base"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, _ := c.Cookie("token")
		data, err := base64.URLEncoding.DecodeString(token)
		if err != nil {
			panic(err)
		}

		token = string(data)

		if token != "" {
			if flag, userId := components.RequireTokenAuthentication(token); flag == true {
				c.Set("userId", userId)
				c.Next()
			}
		}
		c.AbortWithStatusJSON(http.StatusForbidden, base.Fail("authv1 failure"))
	}
}
