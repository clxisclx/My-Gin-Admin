package config

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"path/filepath"
)

func Viper(path string, workDir *string, config *SvcConfig) *viper.Viper {
	// 工作目录
	var mainDir string
	var err error
	if workDir == nil {
		mainDir, err = filepath.Abs(".")
	} else {
		mainDir, err = filepath.Abs(*workDir)
	}
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

	MGA_LOG.Info("配置文件初始化完成",
		zap.String("config_file", path))

	return v
}
