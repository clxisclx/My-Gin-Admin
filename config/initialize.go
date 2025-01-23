package config

// InitAll 统一初始化所有配置
func InitAll(configFile string) {
	// 初始化日志
	InitLogger()

	// 初始化配置文件
	Viper(configFile, nil, &MGA_CONFIG)

	// gorm 初始化
	GormInit()

	// redis 初始化
	RedisInit()
}
