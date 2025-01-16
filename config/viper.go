package config

import (
	"My-Gin-Admin/global"
	"fmt"
	"github.com/spf13/viper"
	"path/filepath"
)

func Viper(path string, config *global.SvcConfig) *viper.Viper {
	// 工作目录
	mainDir, err := filepath.Abs(".")
	if err != nil {
		panic(fmt.Errorf("failed to get main directory: %w", err))
	} else {
		fmt.Println("main directory: ", mainDir)
	}

	// 读取配置文件
	v := viper.New()
	v.AddConfigPath(mainDir)
	v.SetConfigName(path)
	v.SetConfigType("yaml")
	err = v.ReadInConfig() // Find and read the config file
	if err != nil {        // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	// 结构赋值
	err = v.Unmarshal(config)
	if err != nil {
		panic(fmt.Errorf("failed to unmarshal config: %w", err))
	}

	return v
}
