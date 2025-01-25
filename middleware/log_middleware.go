package middleware

import (
	"My-Gin-Admin/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strings"
	"time"
)

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w responseBodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// LoggerMiddleware 日志中间件
func LoggerMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 生成并设置 traceId
		traceId := utils.GenerateUuid()
		c.Set("traceId", traceId)

		// 记录请求开始的时间
		startTime := time.Now()

		// 获取请求参数
		var requestParams string
		method := c.Request.Method
		if method == "GET" {
			requestParams = c.Request.URL.RawQuery
		} else {
			rawData, _ := c.GetRawData()
			if len(rawData) > 0 {
				requestParams = string(rawData)
			}
			// 重新设置请求体，因为GetRawData会清空body
			c.Request.Body = NewBuffer(rawData)
		}

		// 获取请求头
		headers := make(map[string]string)
		for k, v := range c.Request.Header {
			headers[k] = strings.Join(v, ",")
		}
		headerJSON, _ := json.Marshal(headers)

		// 包装 ResponseWriter 以捕获响应内容
		responseBody := &bytes.Buffer{}
		writer := responseBodyWriter{
			ResponseWriter: c.Writer,
			body:           responseBody,
		}
		c.Writer = writer

		// 继续执行后续的中间件和路由处理
		c.Next()

		// 记录请求的相关信息
		duration := time.Since(startTime)
		statusCode := c.Writer.Status()
		path := c.Request.URL.Path
		clientIP := c.ClientIP()

		// 获取响应内容
		responseContent := responseBody.String()
		// 如果响应内容太长，可以截断
		if len(responseContent) > 1000 {
			responseContent = responseContent[:1000] + "..."
		}

		// 构建自定义格式的日志信息
		timeStr := time.Now().Format("2006-01-02 15:04:05")
		logMessages := []string{
			fmt.Sprintf("[%s] TraceID=%s", timeStr, traceId),
			fmt.Sprintf("URL: %s", path),
			fmt.Sprintf("Method: %s", method),
			fmt.Sprintf("Params: %s", requestParams),
			fmt.Sprintf("Response: %s", responseContent),
			fmt.Sprintf("Headers: %s", string(headerJSON)),
			fmt.Sprintf("ClientIP: %s", clientIP),
			fmt.Sprintf("Status: %d", statusCode),
			fmt.Sprintf("Duration: %s", duration.String()),
			"----------------------------------------",
		}

		logMessage := strings.Join(logMessages, "\n")

		// 根据状态码选择日志级别
		if statusCode >= 400 {
			logger.Error(logMessage)
		} else {
			logger.Info(logMessage)
		}
	}
}

// NewBuffer 创建新的buffer用于重新设置请求体
func NewBuffer(data []byte) *Buffer {
	return &Buffer{data: data}
}

type Buffer struct {
	data []byte
}

func (b *Buffer) Read(p []byte) (n int, err error) {
	copy(p, b.data)
	return len(b.data), nil
}

func (b *Buffer) Close() error {
	return nil
}
