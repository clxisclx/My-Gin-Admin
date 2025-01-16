package main

import (
	"My-Gin-Admin/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	routes.UserRoutesInit(router)
	router.Run(":7070")
}
