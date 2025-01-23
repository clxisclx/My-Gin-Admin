package config

import (
	"context"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type Redis struct {
	Addr     string `json:"prefix" yaml:"addr"`
	Port     string `json:"port" yaml:"port"`
	Password string `json:"password" yaml:"password"`
	Db       int    `json:"db" yaml:"db"`
}

func RedisInit() {

	client := redis.NewClient(&redis.Options{
		Addr:     MGA_CONFIG.Redis.Addr + ":" + MGA_CONFIG.Redis.Port,
		Password: MGA_CONFIG.Redis.Password,
		DB:       MGA_CONFIG.Redis.Db,
	})

	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		MGA_LOG.Error("redis连接失败,err:", zap.Error(err))
		panic(err)
		return
	}

	// redis 连接成功
	MGA_LOG.Info("redisl连接成功,pong:", zap.String("pong", pong))
	MGA_Rdb = client
}
