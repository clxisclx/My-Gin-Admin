package user

import (
	"My-Gin-Admin/config"
	"My-Gin-Admin/controller/user/models"
	"My-Gin-Admin/gorm/DO"
	"github.com/jinzhu/copier"
)

type UserService struct{}

func (u *UserService) CreateUser(userReq *models.UserReq) error {
	var userDO DO.User
	copier.Copy(&userDO, userReq)
	err := config.MGA_DB.Create(&userDO).Error
	return err
}
