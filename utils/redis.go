package utils

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),     // 例如 "localhost:6379"
		Password: os.Getenv("REDIS_PASSWORD"), // 無密碼可留空
		DB:       0,
	})
}

func GetRedisCtx() context.Context {
	return context.Background()
}
