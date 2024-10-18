package queue_handler

import (
	"context"
	"launchpad-go-rest/internal/repository"
	"launchpad-go-rest/internal/service"

	"github.com/hibiken/asynq"
)

type QueueHandler interface {
	HandleTestQueue(ctx context.Context, t *asynq.Task) error
	HandleTestQueue2(ctx context.Context, t *asynq.Task) error
}

type handler struct {
	repositories *repository.Repository
	services     *service.Service
}

func New(repositories *repository.Repository, services *service.Service) QueueHandler {
	return &handler{
		repositories: repositories,
		services:     services,
	}
}
