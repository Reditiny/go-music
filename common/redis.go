package common

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go-music/config"
	"log"
)

// Ctx 顶级上下文，不会被取消，也不会超时
// 可以根据需要对这个顶级上下文进行修改、衍生
// 以实现一些更高级的功能，如设置超时、取消信号、跟踪等。
var Ctx = context.Background()
var RedisClient *redis.Client

func InitRedis() {
	// 连接 Redis 数据库
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Conf.Redis.Host, config.Conf.Redis.Port),
		Password: fmt.Sprintf("%s", config.Conf.Redis.Password),
	})
	// 测试连接是否成功
	_, err := RedisClient.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
}
