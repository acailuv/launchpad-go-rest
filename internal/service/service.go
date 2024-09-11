package service

import (
	"launchpad-go-rest/internal/lib/utils"
	"launchpad-go-rest/internal/repository"
	"launchpad-go-rest/internal/service/user"

	"github.com/labstack/echo/v4"
)

type Service struct {
	User user.Service
}

func Init(repositories *repository.Repository, logger echo.Logger, utils utils.Utils) *Service {
	return &Service{
		User: user.New(repositories.User, logger, utils),
	}
}
