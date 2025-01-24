package sys

import (
	"My-Gin-Admin/config"
	"My-Gin-Admin/constants"
	"My-Gin-Admin/gorm/model"
	"My-Gin-Admin/models/req"
	"My-Gin-Admin/utils"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type SysService struct{}

// Login 登录
func (this *SysService) Login(req req.LoginReq) (string, error) {

	var sysUser model.SysUser
	err2 := copier.Copy(&sysUser, req)
	if err2 != nil {
		config.MGA_LOG.Error("复制用户失败")
		return "", err2
	}

	err := config.MGA_DB.Where("username = ?", sysUser.Username).First(&sysUser).Error
	if err != nil {
		return "", errors.New("用户不存在")
	}

	// 校验密码
	if !utils.ComparePassword(sysUser.Password, req.Password, sysUser.Salt) {
		return "", errors.New("密码错误")
	}

	return "temp token", nil // TODO JWT
}

// Register 注册
func (this *SysService) Register(req req.RegisterReq) error {

	var sysUser model.SysUser
	err3 := copier.Copy(&sysUser, req)
	if err3 != nil {
		config.MGA_LOG.Error("复制用户失败")
		return err3
	}
	// 判断用户名是否注册
	err2 := config.MGA_DB.Where("username = ?", sysUser.Username).First(&sysUser).Error
	if !errors.Is(err2, gorm.ErrRecordNotFound) {
		return constants.DB_Data_Exist
	}

	// 随机生成盐
	salt, err := utils.GenerateSalt(16)
	if err != nil {
		config.MGA_LOG.Error("生成盐失败")
		return err
	}
	sysUser.Salt = salt

	// 密码加密
	password := utils.HashPassword(sysUser.Password, []byte(salt))
	sysUser.Password = password
	// 保存
	err = config.MGA_DB.Create(&sysUser).Error
	if err != nil {
		config.MGA_LOG.Error("保存用户失败")
		return err
	}

	return nil
}
