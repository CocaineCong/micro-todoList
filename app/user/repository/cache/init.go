package cache

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"

	"github.com/CocaineCong/micro-todoList/config"
)

// RedisClient Redis缓存客户端单例
var RedisClient *redis.Client

// InitCache 在中间件中初始化redis链接
func InitCache() {
	host := config.RedisHost
	port := config.RedisPort
	password := config.RedisPassword
	database := config.RedisDbName
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
		DB:       database,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	RedisClient = client
}
