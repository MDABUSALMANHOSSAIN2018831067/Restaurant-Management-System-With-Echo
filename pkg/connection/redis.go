package connection

import (
	"context"
	"fmt"
	"restaurant-management/pkg/config"

	"github.com/go-redis/redis/v8"
)

var Client *redis.Client

func RedisConnection() {

	Client = redis.NewClient(&redis.Options{
		Addr:     config.LocalConfig.REDIS_HOST + ":" + config.LocalConfig.REDIS_PORT,
		Password: config.LocalConfig.REDIS_PASS,
		DB:       0,
	})
	red, err := Client.Ping(context.Background()).Result()
	fmt.Println(red)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("redis connection successful...")
}

func Redis() *redis.Client {
	if Client == nil {
		RedisConnection()
	}
	return Client
}
