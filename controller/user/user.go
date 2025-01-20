package user

import (
	"My-Gin-Admin/controller/user/models/req"
	"My-Gin-Admin/models/common/response"
	"github.com/gin-gonic/gin"
)

type UserApi struct{}

func (api UserApi) Index(ctx *gin.Context) {
	response.Success("用户首页", ctx)
}

func (api UserApi) Create(ctx *gin.Context) {
	var user req.UserReq
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

	response.Success(nil, ctx)
}

func (api UserApi) QueryById(ctx *gin.Context) {
	var query req.QueryById
	err := ctx.ShouldBindQuery(&query)
	if err != nil {
		response.Fail(ctx)
		return
	}

	user, err := userService.QueryById(&query)
	if err != nil {
		response.QueryNotData(ctx)
		return
	}

	response.Success(user, ctx)
}

func (api UserApi) Delete(ctx *gin.Context) {
	var query req.QueryById
	err := ctx.ShouldBindQuery(&query)
	if err != nil {
		response.Fail(ctx)
		return
	}

	err = userService.Delete(&query)
	if err != nil {
		response.Fail(ctx)
		return
	}

	response.Success(nil, ctx)
}

func (api UserApi) Update(ctx *gin.Context) {
	var user req.UserReq
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		response.Fail(ctx)
		return
	}

	err = userService.UpdateUser(&user)
	if err != nil {
		response.Fail(ctx)
		return
	}

	response.Success(nil, ctx)
}
