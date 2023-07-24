package database

import (
	"clockwork-server/config"
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type DBRedis interface {
	setValue(context.Context, string, interface{}, int) error
	getValue(context.Context, string) ([]byte, error)
}

type dbRedis struct {
	redisClient redis.Client
}

func NewDBRedis(config *config.Config) DBRedis {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Redis.Host,
		Password: config.Redis.Password,
		DB:       0,
	})

	return &dbRedis{
		redisClient: *rdb,
	}
}

func (d *dbRedis) setValue(ctx context.Context, key string, value interface{}, lifetime int) error {
	return d.redisClient.Set(ctx, key, value, 6*time.Minute).Err()
}

func (d *dbRedis) getValue(ctx context.Context, key string) ([]byte, error) {
	val, err := d.redisClient.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	return []byte(val), nil
}
