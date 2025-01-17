package test

import (
	"My-Gin-Admin/config"
	auto "My-Gin-Admin/gorm"
	"gorm.io/gorm"
	"testing"
	"time"
)

type User struct {
	gorm.Model
	DeletedAt time.Time `gorm:"index:idx_email_deleted_at;index:idx_phone_deleted_at;column:deleted_at;comment:删除时间"`
	Name      string    `gorm:"column:name;type:varchar(32);comment:姓名"`
	Email     string    `gorm:"column:email;type:varchar(64);comment:邮箱;index:idx_email_deleted_at"`
	Password  string    `gorm:"column:password;type:text;comment:密码"`
	Phone     string    `gorm:"column:phone;type:varchar(16);comment:手机号;index:idx_phone_deleted_at"`
	Avatar    string    `gorm:"column:avatar;type:text;comment:头像"`
}

func init() {
	var workDir = "../"
	config.Viper("config.yaml", &workDir, &config.MGA_CONFIG)
	config.GormInit()
}

func TestCreateTableAuto(t *testing.T) {
	db := config.MGA_DB
	auto.AutoTable{}.CreateTableAuto(db, &User{}, "user", "用户表")
}
