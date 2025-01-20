package routes

import (
	"My-Gin-Admin/controller/user"
	"github.com/gin-gonic/gin"
)

func UserRoutesInit(router *gin.Engine) {

	adminRouter := router.Group("/user")
	{
		adminRouter.GET("/index", user.UserApi{}.Index)
		adminRouter.POST("/create", user.UserApi{}.Create)
		adminRouter.POST("/delete", user.UserApi{}.Delete)
		adminRouter.POST("/update", user.UserApi{}.Update)
		adminRouter.GET("/query", user.UserApi{}.QueryById)
	}

}
