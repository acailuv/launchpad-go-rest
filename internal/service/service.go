package service

import (
	"launchpad-go-rest/internal/lib/utils"
	"launchpad-go-rest/internal/repository"
	"launchpad-go-rest/internal/service/user"
	utils_service "launchpad-go-rest/internal/service/utils"
)

type Service struct {
	User  user.Service
	Utils utils_service.Service
}

func Init(repositories *repository.Repository, utils utils.Utils) *Service {
	return &Service{
		User:  user.New(repositories.User, utils, repositories.Cache),
		Utils: utils_service.New(repositories.Publisher),
	}
}
