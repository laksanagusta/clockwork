package database

import (
	"clockwork-server/config"
	"context"

	"github.com/redis/go-redis/v9"
)

type DBRedis interface {
	setValue(context.Context, string, interface{}, int) error
	getValue(context.Context, string) error
}

type dbRedis struct {
	redisConfig config.Redis
	redisClient redis.Client
}

func newDBRedis(redisConfig config.Redis) DBRedis {
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Host,
		Password: redisConfig.Password,
		DB:       0,
	})

	return &dbRedis{
		redisConfig: redisConfig,
		redisClient: *rdb,
	}
}

func (d *dbRedis) setValue(ctx context.Context, key string, value interface{}, lifetime int) error {
	err := d.redisClient.Set(ctx, key, value, 0)
	if err != nil {
		panic(err)
	}

	return nil
}

func (d *dbRedis) getValue(ctx context.Context, key string) error {
	err := d.redisClient.Get(ctx, key)
	if err != nil {
		panic(err)
	}

	return nil
}
