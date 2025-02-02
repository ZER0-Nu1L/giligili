package cache

import (
	"os"
	"strconv"
	"context"

	"github.com/go-redis/redis/v8"
)

// RedisClient Redis缓存客户端单例
var RedisClient *redis.Client
var Ctx = context.Background()

// Redis 在中间件中初始化redis链接
func Redis() {
	db, _ := strconv.ParseUint(os.Getenv("REDIS_DB"), 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PW"),
		DB:       int(db),
	})

	_, err := client.Ping(Ctx).Result()

	if err != nil {
		panic(err)
	}

	RedisClient = client
}
