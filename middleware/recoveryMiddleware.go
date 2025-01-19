package middleware

import (
	"My-Gin-Admin/models/common/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CustomRecovery 自定义 Recovery 中间件
func CustomRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				// 捕获 panic 信息
				c.JSON(http.StatusInternalServerError, response.Response{
					Code: http.StatusInternalServerError,
					Msg:  "网络异常",
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}
