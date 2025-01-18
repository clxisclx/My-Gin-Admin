package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	MGA_VIPER  *viper.Viper
	MGA_CONFIG SvcConfig
	MGA_DB     *gorm.DB
	MGA_LOG    *zap.Logger
)

type SvcConfig struct {
	Name string
	DB   DB
}
