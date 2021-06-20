package utils

import (
	"os"

	"github.com/go-redis/redis"
)

func GetClientRedis() redis.Client {
	opt, _ := redis.ParseURL(os.Getenv("REDIS_URL"))
	client := redis.NewClient(opt)
	return *client
}
