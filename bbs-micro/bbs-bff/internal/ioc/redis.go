package ioc

import "github.com/redis/go-redis/v9"

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-26 11:44

func InitRedis(cfg *Config) redis.Cmdable {
	addr := cfg.Redis.DSN
	redisClient := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	return redisClient
}
