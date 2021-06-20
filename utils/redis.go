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
		var local = false
		if local {
			client = redis.NewClient(&redis.Options{
				Addr:     "localhost:6379",
				Password: "", // no password set
				DB:       0,  // use default DB
			})
			return *client
		} else {
			opt, _ := redis.ParseURL(os.Getenv("REDIS_URL"))
			client = redis.NewClient(opt)
			return *client
		}
	}
	return *client
}
