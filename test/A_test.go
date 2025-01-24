package test

import "My-Gin-Admin/config"

func init() {
	config.InitLogger()
	var workDir = "../"
	config.Viper("config_clx.yaml", &workDir, &config.MGA_CONFIG)
}
