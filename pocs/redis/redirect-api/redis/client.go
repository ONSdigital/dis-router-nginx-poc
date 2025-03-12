package redis

import (
	redis "github.com/redis/go-redis/v9"
)

func NewRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:             "redis:6379",
		Password:         "",
		DB:               0,
		DisableIndentity: true,
	})
	return client
}
