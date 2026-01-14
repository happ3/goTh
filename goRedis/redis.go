package goRedis

import "github.com/redis/go-redis/v9"

var RDB *redis.Client

func InitRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123456Admin@123", // 根据实际配置
		DB:       0,
	})
}
