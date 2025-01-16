package config

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	MGA_VIPER  *viper.Viper
	MGA_CONFIG SvcConfig
	MGA_DB     *gorm.DB
)

type SvcConfig struct {
	Name string
	DB   DB
}
