package user

import (
	"context"
	"database/sql"
	"launchpad-go-rest/internal/lib/errors"
	"launchpad-go-rest/internal/repository/cache"
	"launchpad-go-rest/pkg/types/user"
	"net/http"
	"time"

	"github.com/go-redis/redis"
	"github.com/rs/zerolog/log"
)

func (s service) FindByID(ctx context.Context, p user.FindByIDRequest) (user.FindByIDResponse, error) {
	if err := p.Validate(ctx); err != nil {
		return user.FindByIDResponse{}, errors.NewWithCode(http.StatusBadRequest, errors.VALIDATION_ERROR, err.Error())
	}

	var res user.FindByIDResponse
	err := s.cache.Get(ctx, cache.FindUserByIDCacheKey(p.ID), &res)
	if err != nil && !errors.Is(err, redis.Nil) {
		log.Error().Stack().Err(err).Msg("Redis error")
	} else if err == nil {
		return res, nil
	}

	user, err := s.user.FindByID(ctx, p.ID)
	if errors.Is(err, sql.ErrNoRows) {
		return user.ToFindByIDResponse(), errors.NewWithCode(http.StatusBadRequest, errors.USER_NOT_FOUND, "Invalid user id")
	} else if err != nil {
		return user.ToFindByIDResponse(), err
	}

	err = s.cache.Set(ctx, cache.FindUserByIDCacheKey(p.ID), user.ToFindByIDResponse(), time.Minute)
	if err != nil {
		log.Error().Stack().Err(err).Msg("Redis error")
	}

	return user.ToFindByIDResponse(), nil
}
