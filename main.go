package main

import (
	"My-Gin-Admin/config"
	"My-Gin-Admin/middleware"
	"My-Gin-Admin/routes"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	// 统一初始化配置
	config.InitAll("config.yaml")
	// 路由
	router := gin.Default()
	// 日志中间件
	router.Use(middleware.LoggerMiddleware(config.MGA_LOG))

	routes.UserRoutesInit(router)
	config.MGA_LOG.Info("启动服务器", zap.String("port", ":7070"))
	router.Run(":7070")
}
