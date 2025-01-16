package routes

import (
	"My-Gin-Admin/controller/user"
	"github.com/gin-gonic/gin"
)

func UserRoutesInit(router *gin.Engine) {

	adminRouter := router.Group("/user")
	{
		adminRouter.GET("/index", user.UserController{}.Index)
	}

}
