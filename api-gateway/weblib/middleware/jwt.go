package middleware

import (
	"api-gateway/pkg/utils"
	"github.com/gin-gonic/gin"
)

// JWT token验证中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code uint32

		code = 200
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 404
		} else {
			_, err := utils.ParseToken(token)
			if err != nil {
				code = 401
			}
		}
		if code != 200 {
			c.JSON(500, gin.H{
				"code": code,
				"msg":  "鉴权失败",
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
