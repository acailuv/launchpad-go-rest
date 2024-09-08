package service

import (
	"launchpad-go-rest/internal/repository"
	"launchpad-go-rest/internal/service/user"

	"github.com/labstack/echo/v4"
)

type Service struct {
	User user.Service
}

func Init(repositories *repository.Repository, logger echo.Logger) *Service {
	return &Service{
		User: user.New(repositories.User, logger),
	}
}
