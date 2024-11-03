package utils

import (
	"os"
	"sync"
	"github.com/go-redis/redis"
)

var lock = &sync.Mutex{}
var client *redis.Client

func GetClientRedis() redis.Client {
	if client == nil {
		lock.Lock()
		defer lock.Unlock()
		// local redis: "localhost:6379"
		var urlRedis = os.Getenv("REDIS_URL")
		client = redis.NewClient(&redis.Options{
			Addr:     urlRedis,
			Password: "", // no password set
			DB:       0,  // use default DB
		})
		return *client
	}
	return *client
}
