package config

import "go.uber.org/zap"

// InitAll 统一初始化所有配置
func InitAll(configFile string) {
	// 初始化日志
	InitLogger()

	// 初始化配置文件
	Viper(configFile, nil, &MGA_CONFIG)

	MGA_LOG.Info("配置文件初始化完成",
		zap.String("config_file", configFile))

	MGA_LOG.Info("数据库配置初始化完成",
		zap.String("path", MGA_CONFIG.DB.Mysql.Path),
		zap.String("port", MGA_CONFIG.DB.Mysql.Port),
		zap.String("dbname", MGA_CONFIG.DB.Mysql.Dbname))

	// gorm 初始化
	GormInit()
	MGA_LOG.Info("gorm初始化成功")
}
