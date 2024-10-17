package user

import (
	"context"
	"launchpad-go-rest/internal/lib/utils"
	"launchpad-go-rest/internal/repository/cache"
	"launchpad-go-rest/internal/repository/user"
	user_types "launchpad-go-rest/pkg/types/user"
)

type Service interface {
	Find(ctx context.Context) ([]user_types.FindResponse, error)
	FindByID(ctx context.Context, p user_types.FindByIDRequest) (user_types.FindByIDResponse, error)
	Create(ctx context.Context, p user_types.CreateRequest) error
	UpdateByID(ctx context.Context, p user_types.UpdateByIDRequest) error
	DeleteByID(ctx context.Context, p user_types.DeleteByIDRequest) error
}

type service struct {
	user  user.Repository
	utils utils.Utils
	cache cache.Repository
}

func New(
	user user.Repository,
	utils utils.Utils,
	cache cache.Repository,
) Service {
	return &service{
		user:  user,
		utils: utils,
		cache: cache,
	}
}
