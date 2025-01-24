package routes

import (
	"My-Gin-Admin/controller/sys"
	"github.com/gin-gonic/gin"
)

func SysRoutesInit(router *gin.Engine) {

	r := router.Group("/sys")
	{
		r.POST("/login", sys.SysApi{}.Login)
		r.POST("/register", sys.SysApi{}.Register)
	}
}
