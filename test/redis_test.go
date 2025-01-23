package test

import (
	"My-Gin-Admin/config"
	"context"
	"testing"
)

func init() {
	var workDir = "../"
	config.Viper("config.yaml", &workDir, &config.MGA_CONFIG)
	config.RedisInit()
}

func TestRedis(t *testing.T) {

	result, err := config.MGA_Rdb.Get(context.Background(), "k1").Result()
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}
