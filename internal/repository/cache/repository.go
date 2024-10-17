package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis"
)

type Repository interface {
	Get(ctx context.Context, key string, target any) error
	Set(ctx context.Context, key string, value any, exp time.Duration) error
}

type repository struct {
	redis *redis.Client
}

func New(redis *redis.Client) Repository {
	return &repository{
		redis: redis,
	}
}
