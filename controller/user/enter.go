package user

import "My-Gin-Admin/service"

type ApiGroup struct{ UserApi }

var userService = service.ServiceGroupApp.UserServiceGroup.UserService
