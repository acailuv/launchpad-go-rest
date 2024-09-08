package user

import (
	"context"
	"launchpad-go-rest/pkg/types/user"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type Repository interface {
	Find(ctx context.Context) ([]user.User, error)
	FindByID(ctx context.Context, id string) (user.User, error)
	Create(ctx context.Context, u user.User) error
	UpdateByID(ctx context.Context, u user.User) error
	DeleteByID(ctx context.Context, id string) error
}

type repository struct {
	db     *sqlx.DB
	logger echo.Logger
}

func New(db *sqlx.DB, logger echo.Logger) Repository {
	return &repository{
		db:     db,
		logger: logger,
	}
}
