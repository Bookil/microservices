package db

import (
	"fmt"
	"sync"

	"github.com/Bookil/microservices/auth/config"
	"github.com/go-redis/redis/v8"
)

var (
	redisLock     = &sync.Mutex{}
	redisInstance *redis.Client
)

func GetRedisInstance(config config.Redis) *redis.Client {
	if redisInstance == nil {
		redisLock.Lock()
		defer redisLock.Unlock()

		client := redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
			Password: config.Password,
			DB:       config.DB,
		})

		redisInstance = client
	}
	return redisInstance
}
