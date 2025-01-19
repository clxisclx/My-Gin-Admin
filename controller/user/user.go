package user

import (
	"My-Gin-Admin/controller/user/models"
	"My-Gin-Admin/models/common/response"
	"github.com/gin-gonic/gin"
)

type UserApi struct{}

func (c UserApi) Index(ctx *gin.Context) {
	response.Success("用户首页", ctx)
}

func (c UserApi) Create(ctx *gin.Context) {

	var user models.UserReq
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		response.Fail(ctx)
		return
	}

	err = userService.CreateUser(&user)
	if err != nil {
		response.Fail(ctx)
		return
	}

	response.Success("创建用户成功", ctx)

}
