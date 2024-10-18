package utils

import (
	"context"
	"launchpad-go-rest/internal/repository/publisher"
	utils_type "launchpad-go-rest/pkg/types/utils"
)

type Service interface {
	PublishTask(ctx context.Context, p utils_type.PublishTaskRequest) error
}

type service struct {
	publisher publisher.Repository
}

func New(
	publisher publisher.Repository,
) Service {
	return &service{
		publisher: publisher,
	}
}
