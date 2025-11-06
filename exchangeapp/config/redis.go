package config

import (
	"log"

	"github.com/go-redis/redis"

	"exchangeapp/global"
)

func initRedis() {
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     AppConfig.Redis.Address,
		DB:       AppConfig.Redis.DB,
		Password: AppConfig.Redis.password,
	})
	_, err := RedisClient.Ping().Result()
	if err != nil {
		log.Fatalf("Failed to connect to redis: %v", err)
	}
	global.RedisDB = RedisClient
}
