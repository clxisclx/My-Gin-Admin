package sys

import (
	"My-Gin-Admin/constants"
	"My-Gin-Admin/models/common/response"
	"My-Gin-Admin/models/req"
	"errors"
	"github.com/gin-gonic/gin"
)

type SysApi struct{}

// Login 登录
func (api SysApi) Login(ctx *gin.Context) {
	var req req.LoginReq
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		response.ParamError(ctx)
		return
	}

	token, err := sysService.Login(req)
	if err != nil {
		response.LoginFail(ctx, err.Error())
		return
	}
	response.Success(token, ctx)
}

// 注册
func (api SysApi) Register(ctx *gin.Context) {
	var req req.RegisterReq
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		response.ParamError(ctx)
		return
	}

	// 校验用户名长度
	username := req.Username
	if len(username) < 4 || len(username) > 20 {
		response.FailWithMsg(ctx, "用户名长度为4-20位")
		return
	}

	// 密码校验长度
	password := req.Password
	if len(password) < 6 || len(password) > 20 {
		response.FailWithMsg(ctx, "密码长度为6-20位")
		return
	}

	err = sysService.Register(req)
	if errors.Is(err, constants.DB_Data_Exist) {
		response.FailWithMsg(ctx, "用户已存在")
		return
	}

	if err != nil {
		response.FailWithMsg(ctx, "注册用户失败")
		return
	}
	response.Success(nil, ctx)
}
