package test

import (
	"My-Gin-Admin/config"
	auto "My-Gin-Admin/gorm"
	"My-Gin-Admin/gorm/model"
	"gorm.io/gorm"
	"testing"
)

var db *gorm.DB

func init() {
	config.GormInit()
	db = config.MGA_DB
}

func TestCreateTableAuto(t *testing.T) {
	auto.AutoTable{}.CreateTableAuto(db, &model.User{}, "user", "用户表")
}

func TestCreateSysUserTableAuto(t *testing.T) {
	auto.AutoTable{}.CreateTableAuto(db, &model.SysUser{}, "sys_user", "系统用户表")
}
