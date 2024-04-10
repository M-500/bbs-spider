package dep_setup

import (
	"context"
	"github.com/redis/go-redis/v9"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-10 17:20

var redisClient redis.Cmdable

func InitRedis() redis.Cmdable {
	if redisClient == nil {
		redisClient = redis.NewClient(&redis.Options{
			Addr: "192.168.1.52:6379",
		})

		for err := redisClient.Ping(context.Background()).Err(); err != nil; {
			panic(err)
		}
	}
	return redisClient
}
