package repository

import (
	"launchpad-go-rest/internal/repository/user"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	User user.Repository
}

func Init(db *sqlx.DB) *Repository {
	return &Repository{
		User: user.New(db),
	}
}
