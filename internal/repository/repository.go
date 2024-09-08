package repository

import (
	"launchpad-go-rest/internal/repository/user"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type Repository struct {
	User user.Repository
}

func Init(db *sqlx.DB, logger echo.Logger) *Repository {
	return &Repository{
		User: user.New(db, logger),
	}
}
