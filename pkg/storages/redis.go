package storages

import (
	"context"
	"ecm-api-template/internal/configs"
	"log"

	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

func GetRedis() *redis.Client {
	if redisClient == nil {
		log.Fatal("database_not_initialized")
	}
	return redisClient
}

func NewRedis() {
	redisOpts, err := redis.ParseURL(configs.Environment.REDIS_URI)
	if err != nil {
		log.Fatal(err)
	}
	redisClient = redis.NewClient(redisOpts)

	// Check connection
	cmd := redisClient.Ping(context.Background())
	if cmd.Err() != nil {
		log.Fatal("Redis ping failed ", cmd.Err())
	}
	log.Println("Redis is connected,", configs.Environment.REDIS_URI)
}
