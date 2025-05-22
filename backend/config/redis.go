package config

import (
	"os"

	"github.com/redis/go-redis/v9"
)

var Client *redis.Client

func InitRedisClient() {

	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Username: os.Getenv("REDIS_SESSION_USERNAME"),
		Password: os.Getenv("REDIS_SESSION_PASSWORD"),
		Protocol: 2,
	})

	Client = client
}
