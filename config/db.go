package config

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	Mysql Mysql `json:"mysql" yaml:"mysql"`
}

type Mysql struct {
	Prefix       string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                         // 数据库前缀
	Port         string `mapstructure:"port" json:"port" yaml:"port"`                               // 数据库端口
	Config       string `mapstructure:"config" json:"config" yaml:"config"`                         // 高级配置
	Dbname       string `mapstructure:"db-name" json:"db-name" yaml:"db-name"`                      // 数据库名
	Username     string `mapstructure:"username" json:"username" yaml:"username"`                   // 数据库账号
	Password     string `mapstructure:"password" json:"password" yaml:"password"`                   // 数据库密码
	Path         string `mapstructure:"path" json:"path" yaml:"path"`                               // 数据库地址
	Engine       string `mapstructure:"engine" json:"engine" yaml:"engine" default:"InnoDB"`        // 数据库引擎，默认InnoDB
	LogMode      string `mapstructure:"log-mode" json:"log-mode" yaml:"log-mode"`                   // 是否开启Gorm全局日志
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns"` // 空闲中的最大连接数
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"max-open-conns" yaml:"max-open-conns"` // 打开到数据库的最大连接数
	Singular     bool   `mapstructure:"singular" json:"singular" yaml:"singular"`                   // 是否开启全局禁用复数，true表示开启
	LogZap       bool   `mapstructure:"log-zap" json:"log-zap" yaml:"log-zap"`                      // 是否通过zap写入日志文件
}

func Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", MGA_CONFIG.DB.Mysql.Username, MGA_CONFIG.DB.Mysql.Password, MGA_CONFIG.DB.Mysql.Path, MGA_CONFIG.DB.Mysql.Port, MGA_CONFIG.DB.Mysql.Dbname, MGA_CONFIG.DB.Mysql.Config)
}

func GormInit() {
	init, err := gorm.Open(mysql.Open(Dsn()), &gorm.Config{
		Logger: logger.New(
			zap.NewStdLog(MGA_LOG), // 使用 zap 日志
			logger.Config{
				LogLevel: logger.Info, // 设置 GORM 日志等级为 Info
				Colorful: false,       // 启用彩色日志
			},
		)})
	if err != nil {
		MGA_LOG.Error("gorm初始化失败", zap.Error(err))
		panic(err)
	}
	MGA_DB = init

	MGA_LOG.Info("数据库配置初始化完成",
		zap.String("path", MGA_CONFIG.DB.Mysql.Path),
		zap.String("port", MGA_CONFIG.DB.Mysql.Port),
		zap.String("dbname", MGA_CONFIG.DB.Mysql.Dbname))

	MGA_LOG.Info("gorm初始化成功")

}
