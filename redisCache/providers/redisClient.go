package providers

import (
	"context"
	"log"
	"time"

	redis "github.com/go-redis/redis/v8"
)

type CacheClient struct {
	Client *redis.Client
	Ctx    context.Context
}

func NewCacheClient() CacheProvider {

	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6380",
	})

	err := client.Ping(context.Background()).Err()
	if err != nil {
		log.Fatalf("unable to connect redis server %v", err)
	}

	return &CacheClient{
		Client: client,
		Ctx:    context.Background(),
	}
}

func (c *CacheClient) Get(Key string) string {
	return c.Client.Get(c.Ctx, Key).Val()
}

func (c *CacheClient) Set(key string, value interface{}, expiryTime time.Duration) error {
	return c.Client.Set(c.Ctx, key, value, expiryTime).Err()
}

func (c *CacheClient) Delete(key string) int64 {
	return c.Client.Del(c.Ctx, key).Val()
}
