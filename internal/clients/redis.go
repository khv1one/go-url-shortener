package clients

import (
	"time"

	"github.com/go-redis/redis"
)

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

func (r RedisClient) SetIfNotExist(key string, value interface{}, ttl time.Duration) error {
	_, err := r.SetNX(key, value, ttl).Result()

	return err
}

func (r RedisClient) GetByKey(key string) (string, error) {
	return r.Get(key).Result()
}
