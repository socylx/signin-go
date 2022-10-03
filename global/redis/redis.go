package redis

import (
	"context"
	"fmt"
	"log"
	"signin-go/global/config"
	"signin-go/global/time"

	"github.com/go-redis/redis/v8"
)

type Client = redis.Client

var Redis *Client

func Init() {
	log.Println("global.redis.Init Start...")
	Redis := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%v:%v", config.Redis.Host, config.Redis.Port),
		Password:     config.Redis.Password,
		DB:           config.Redis.DB,
		MaxRetries:   3,
		PoolSize:     10,
		MinIdleConns: 5,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := Redis.Ping(ctx).Err(); err != nil {
		log.Fatalf("global.redis.Init.Ping() Error: %v", err)
	}
}

func Close() {
	err := Redis.Close()
	if err != nil {
		log.Printf("global.redis.Close Error: %v", err)
	}
}
