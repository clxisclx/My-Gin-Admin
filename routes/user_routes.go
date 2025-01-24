package routes

import (
	"My-Gin-Admin/controller/user"
	"github.com/gin-gonic/gin"
)

func UserRoutesInit(router *gin.Engine) {

	r := router.Group("/user")
	{
		r.GET("/index", user.UserApi{}.Index)
		r.POST("/create", user.UserApi{}.Create)
		r.POST("/delete", user.UserApi{}.Delete)
		r.POST("/update", user.UserApi{}.Update)
		r.GET("/query", user.UserApi{}.QueryById)
	}

}
