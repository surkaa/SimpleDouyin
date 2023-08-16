package db

import (
	"SimpleDouyin/config"
	"github.com/go-redis/redis/v8"
)

// GetRedisConnection 获取redis连接
func GetRedisConnection() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     config.Config.Redis.Addr,
		Password: config.Config.Redis.Password,
		DB:       config.Config.Redis.DB,
		PoolSize: config.Config.Redis.PoolSize,
	})
}
