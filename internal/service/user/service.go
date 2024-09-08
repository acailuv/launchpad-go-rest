package user

import (
	"context"
	"launchpad-go-rest/internal/repository/user"
	user_types "launchpad-go-rest/pkg/types/user"

	"github.com/labstack/echo/v4"
)

type Service interface {
	Find(ctx context.Context) ([]user_types.FindResponse, error)
	FindByID(ctx context.Context, id string) (user_types.FindByIDResponse, error)
	Create(ctx context.Context, u user_types.CreateRequest) error
	UpdateByID(ctx context.Context, u user_types.UpdateByIDRequest) error
	DeleteByID(ctx context.Context, id string) error
}

type service struct {
	user   user.Repository
	logger echo.Logger
}

func New(
	user user.Repository,
	logger echo.Logger,
) Service {
	return &service{
		user:   user,
		logger: logger,
	}
}
