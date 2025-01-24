package controller

import (
	"My-Gin-Admin/controller/sys"
	"My-Gin-Admin/controller/user"
)

var ApiGroupApp = new(ApiGroup)

//DB := config.MGA_DB

type ApiGroup struct {
	SysApiGroup  sys.ApiGroup
	UserApiGroup user.ApiGroup
}
