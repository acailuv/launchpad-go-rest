package cache

import (
	"context"
	"launchpad-go-rest/internal/lib/errors"
	"time"
)

func (r repository) Set(ctx context.Context, key string, value any, exp time.Duration) error {
	err := r.redis.Set(key, value, exp).Err()
	if err != nil {
		return errors.New(err)
	}

	return nil
}
