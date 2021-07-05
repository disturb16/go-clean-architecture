package rediscache

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/disturb16/go-sqlite-service/settings"
	"github.com/go-redis/redis/v8"
)

type Cache struct {
	client     *redis.Client
	expiration time.Duration
	enabled    bool
}

func New(ctx context.Context, config *settings.Settings) (*Cache, error) {
	opts := &redis.Options{
		Addr: fmt.Sprintf("%s:%d", config.Redis.Address, config.Redis.Port),
		DB:   config.Redis.DBIndex,
	}

	client := redis.NewClient(opts)
	err := client.Ping(ctx).Err()
	if err != nil {
		return nil, err
	}

	return &Cache{
		client:     client,
		expiration: config.Redis.ExpirationMinutes * time.Second,
		enabled:    config.Service.CacheEnabled,
	}, nil
}

func (rc Cache) Set(ctx context.Context, key string, value interface{}) error {
	if !rc.enabled {
		return nil
	}

	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return rc.client.Set(ctx, key, bytes, rc.expiration).Err()
}

func (rc Cache) Get(ctx context.Context, key string, v interface{}) error {
	if !rc.enabled {
		return nil
	}

	val, err := rc.client.Get(ctx, key).Result()
	if err != nil {
		log.Println(err)
	}

	return json.Unmarshal([]byte(val), v)
}

func (rc Cache) Delete(ctx context.Context, key string) error {
	if !rc.enabled {
		return nil
	}

	return rc.client.Del(ctx, key).Err()
}

func (rc Cache) Close() error {
	if !rc.enabled {
		return nil
	}

	return rc.client.Close()
}
