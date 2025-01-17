package main

import (
	"My-Gin-Admin/config"
	"My-Gin-Admin/routes"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	// 初始化配置文件
	_ = config.Viper("config.yaml", nil, &config.MGA_CONFIG)

	fmt.Printf("配置文件初始化完成,%s\n", "config.yaml")
	fmt.Printf("数据库配置初始化完成,%s:%s/%s\n", config.MGA_CONFIG.DB.Mysql.Path, config.MGA_CONFIG.DB.Mysql.Port, config.MGA_CONFIG.DB.Mysql.Dbname)

	// gorm 初始化
	config.GormInit()
	fmt.Println("gorm初始化成功")

	router := gin.Default()
	routes.UserRoutesInit(router)
	router.Run(":7070")
}
