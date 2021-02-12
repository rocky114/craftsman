package cache

import (
	"craftsman/config"
	"github.com/go-redis/redis/v8"
)

var RedisConn *redis.Client

func Bootstrap() {
	redisConfig := config.GlobalConfig.Cache.Redis

	RedisConn = redis.NewClient(&redis.Options{
		Addr:     redisConfig.Host + ":" + redisConfig.Port,
		Password: redisConfig.Password,
		DB:       redisConfig.Db,
	})
}
