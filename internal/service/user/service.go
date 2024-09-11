package user

import (
	"context"
	"launchpad-go-rest/internal/lib/utils"
	"launchpad-go-rest/internal/repository/user"
	user_types "launchpad-go-rest/pkg/types/user"

	"github.com/labstack/echo/v4"
)

type Service interface {
	Find(ctx context.Context) ([]user_types.FindResponse, error)
	FindByID(ctx context.Context, p user_types.FindByIDRequest) (user_types.FindByIDResponse, error)
	Create(ctx context.Context, p user_types.CreateRequest) error
	UpdateByID(ctx context.Context, p user_types.UpdateByIDRequest) error
	DeleteByID(ctx context.Context, p user_types.DeleteByIDRequest) error
}

type service struct {
	user   user.Repository
	logger echo.Logger
	utils  utils.Utils
}

func New(
	user user.Repository,
	logger echo.Logger,
	utils utils.Utils,
) Service {
	return &service{
		user:   user,
		logger: logger,
		utils:  utils,
	}
}
