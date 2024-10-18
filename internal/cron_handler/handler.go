package cron_handler

import (
	"launchpad-go-rest/internal/repository"
	"launchpad-go-rest/internal/service"
)

type CronHandler interface {
	Heartbeat()
}

type handler struct {
	repositories repository.Repository
	services     service.Service
}

func New(repositories repository.Repository, services service.Service) CronHandler {
	return &handler{
		repositories: repositories,
		services:     services,
	}
}
