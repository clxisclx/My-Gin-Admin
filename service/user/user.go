package user

import (
	"My-Gin-Admin/config"
	"My-Gin-Admin/controller/user/models/req"
	"My-Gin-Admin/controller/user/models/resp"
	"My-Gin-Admin/gorm/DO"
	"github.com/jinzhu/copier"
)

type UserService struct{}

func (svc *UserService) CreateUser(userReq *req.UserReq) error {
	var userDO DO.User
	copier.Copy(&userDO, userReq)
	err := config.MGA_DB.Create(&userDO).Error
	return err
}

func (svc *UserService) QueryById(req *req.QueryById) (*resp.UserResp, error) {
	var userDO DO.User
	err := config.MGA_DB.Where("id = ?", req.ID).First(&userDO).Error
	if err != nil {
		return nil, err
	}
	var userResp resp.UserResp
	copier.Copy(&userResp, &userDO)
	return &userResp, nil
}

func (svc *UserService) Delete(req *req.QueryById) error {
	err := config.MGA_DB.Where("id = ?", req.ID).Delete(&DO.User{}).Error
	return err
}

func (svc *UserService) UpdateUser(userReq *req.UserReq) error {
	var userDO DO.User
	copier.Copy(&userDO, userReq)
	err := config.MGA_DB.Updates(&userDO).Error
	return err
}
