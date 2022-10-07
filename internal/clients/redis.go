package clients

import "github.com/go-redis/redis"

type RedisClient struct {
	*redis.Client
}

func NewRedisClient() RedisClient {
	return RedisClient{redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})}
}
