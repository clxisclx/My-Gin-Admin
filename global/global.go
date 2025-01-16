package global

import (
	"github.com/spf13/viper"
)

var (
	MGA_VIPER  *viper.Viper
	MGA_CONFIG SvcConfig
)

type SvcConfig struct {
	Name string
}
