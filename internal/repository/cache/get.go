package cache

import (
	"context"
	"launchpad-go-rest/internal/lib/errors"
)

func (r repository) Get(ctx context.Context, key string, target any) error {
	err := r.redis.Get(key).Scan(target)
	if err != nil {
		return errors.New(err)
	}

	return nil
}
