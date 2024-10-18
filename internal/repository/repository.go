package repository

import (
	"launchpad-go-rest/internal/repository/cache"
	"launchpad-go-rest/internal/repository/publisher"
	"launchpad-go-rest/internal/repository/user"

	"github.com/go-redis/redis"
	"github.com/hibiken/asynq"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	User      user.Repository
	Cache     cache.Repository
	Publisher publisher.Repository
}

func Init(db *sqlx.DB, redis *redis.Client, asynqClient *asynq.Client) *Repository {
	return &Repository{
		User:      user.New(db),
		Cache:     cache.New(redis),
		Publisher: publisher.New(asynqClient),
	}
}
