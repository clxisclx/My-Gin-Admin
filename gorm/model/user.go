package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	DeletedAt time.Time `gorm:"index:idx_email_deleted_at;index:idx_phone_deleted_at;column:deleted_at;comment:删除时间;DEFAULT:NULL"`
	Name      string    `gorm:"column:name;type:varchar(32);comment:姓名"`
	Email     string    `gorm:"column:email;type:varchar(64);comment:邮箱;index:idx_email_deleted_at"`
	Password  string    `gorm:"column:password;type:text;comment:密码"`
	Phone     string    `gorm:"column:phone;type:varchar(16);comment:手机号;index:idx_phone_deleted_at"`
	Avatar    string    `gorm:"column:avatar;type:text;comment:头像"`
}

func (this User) TableName() string {
	return "user"
}
