package repository

import (
	"launchpad-go-rest/internal/repository/cache"
	"launchpad-go-rest/internal/repository/user"

	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	User  user.Repository
	Cache cache.Repository
}

func Init(db *sqlx.DB, redis *redis.Client) *Repository {
	return &Repository{
		User:  user.New(db),
		Cache: cache.New(redis),
	}
}
