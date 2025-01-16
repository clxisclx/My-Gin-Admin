package user

import "github.com/gin-gonic/gin"

type UserController struct{}

func (c UserController) Index(ctx *gin.Context) {
	ctx.String(200, "用户首页")
}
