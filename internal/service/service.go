package service

import (
	"launchpad-go-rest/internal/lib/utils"
	"launchpad-go-rest/internal/repository"
	"launchpad-go-rest/internal/service/user"
)

type Service struct {
	User user.Service
}

func Init(repositories *repository.Repository, utils utils.Utils) *Service {
	return &Service{
		User: user.New(repositories.User, utils),
	}
}
