package goRedis

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var RDB *redis.Client

func InitRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123456Admin@123", // 根据实际配置
		DB:       0,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := RDB.Ping(ctx).Result()
	if err != nil {
		RDB = nil
		panic("Redis 连接失败: %w" + err.Error())
	} else {
		log.Println("✅Redis 连接成功")
	}

}
