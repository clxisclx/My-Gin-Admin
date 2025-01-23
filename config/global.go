package config

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	MGA_VIPER  *viper.Viper
	MGA_CONFIG SvcConfig
	MGA_DB     *gorm.DB
	MGA_Rdb    redis.UniversalClient
	MGA_LOG    *zap.Logger
)

type SvcConfig struct {
	Name  string
	DB    DB
	Redis Redis
}
