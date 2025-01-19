package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"time"
)

// 日志中间件
func LoggerMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 记录请求开始的时间
		startTime := time.Now()

		// 继续执行后续的中间件和路由处理
		c.Next()

		// 记录请求的相关信息
		duration := time.Since(startTime).Milliseconds()
		statusCode := c.Writer.Status()
		method := c.Request.Method
		path := c.Request.URL.Path
		clientIP := c.ClientIP()

		// 构建日志信息
		logger.Info("Request",
			zap.String("method", method),
			zap.String("path", path),
			zap.String("client_ip", clientIP),
			zap.Int("status_code", statusCode),
			zap.String("duration", strconv.FormatInt(duration, 10)+"ms"),
		)
	}
}
