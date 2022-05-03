package platform

import (
	"appslab-ke/kipchoge-go/configs"
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"log"
)

type Cache struct {
	client *redis.Client
}

func (c Cache) Ping() {
	cmd := c.client.Ping(context.Background())
	log.Println(cmd.String())
}

func (c Cache) Set(key, data string) error {
	panic("implement me")
}

func NewCache() (*Cache, error) {
	options := configs.GetRedisConfig()
	if options == nil {
		return nil, errors.New("failed to get redis config")
	}
	redisClient := redis.NewClient(options)

	return &Cache{
		client: redisClient,
	}, nil
}
