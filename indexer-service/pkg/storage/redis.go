// Package storage provides
package storage

import (
	"fmt"

	"portto_interview/indexer-service/pkg/config"

	"github.com/go-redis/redis/v8"
)

// Cache .
type Cache interface {
	redis.Cmdable

	// add custom methods at here
}

// redisClient .
type redisClient struct {
	*redis.Client
}

// NewRedis .
func NewRedis(cfg *config.RedisConfig) (Cache, error) {
	if len(cfg.Addresses) <= 0 {
		return nil, fmt.Errorf("redis config address is empty")
	}

	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Addresses[0],
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	fmt.Println(client)

	return redisClient{Client: client}, nil
}
