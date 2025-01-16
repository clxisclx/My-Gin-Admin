package main

import (
	"My-Gin-Admin/config"
	"My-Gin-Admin/global"
	"My-Gin-Admin/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	// 初始化配置文件
	_ = config.Viper("config.yaml", &global.MGA_CONFIG)

	println(global.MGA_CONFIG.Name)

	router := gin.Default()
	routes.UserRoutesInit(router)
	router.Run(":7070")
}
