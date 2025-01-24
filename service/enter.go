package service

import (
	"My-Gin-Admin/service/sys"
	"My-Gin-Admin/service/user"
)

type ServiceGroup struct {
	UserServiceGroup user.ServiceGroup
	SysServiceGroup  sys.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
