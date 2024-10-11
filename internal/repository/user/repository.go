package user

import (
	"context"
	"launchpad-go-rest/pkg/types/user"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	Find(ctx context.Context) ([]user.User, error)
	FindByID(ctx context.Context, id string) (user.User, error)
	FindByEmail(ctx context.Context, email string) (user.User, error)
	Create(ctx context.Context, u user.User) error
	UpdateByID(ctx context.Context, u user.User) error
	DeleteByID(ctx context.Context, id string) error
}

type repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return &repository{
		db: db,
	}
}
