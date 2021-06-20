package utils

import (
	"os"

	"github.com/go-redis/redis"
)

func GetClientRedis() redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: "",
		DB:       0,
	})

	return *client
}
