package service

import "My-Gin-Admin/service/user"

type ServiceGroup struct {
	UserServiceGroup user.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
