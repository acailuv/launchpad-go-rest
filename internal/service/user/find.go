package user

import (
	"context"
	"launchpad-go-rest/internal/lib/errors"
	"launchpad-go-rest/internal/repository/cache"
	"launchpad-go-rest/pkg/types/user"
	"time"

	"github.com/go-redis/redis"
	"github.com/rs/zerolog/log"
)

func (s service) Find(ctx context.Context) ([]user.FindResponse, error) {
	var response user.FindResponseList = make([]user.FindResponse, 0)

	err := s.cache.Get(ctx, cache.FindUsersCacheKey, &response)
	if err != nil && !errors.Is(err, redis.Nil) {
		log.Error().Stack().Err(err).Msg("Redis error")
	} else if err == nil {
		return response, nil
	}

	res, err := s.user.Find(ctx)
	if err != nil {
		return response, err
	}

	for _, user := range res {
		response = append(response, user.ToFindResponse())
	}

	err = s.cache.Set(ctx, cache.FindUsersCacheKey, response, time.Minute)
	if err != nil {
		log.Error().Stack().Err(err).Msg("Redis error")
	}

	return response, nil
}
