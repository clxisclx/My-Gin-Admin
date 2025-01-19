package user

import (
	"My-Gin-Admin/models/common/response"
	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (c UserController) Index(ctx *gin.Context) {
	response.Success("用户首页", ctx)
}
