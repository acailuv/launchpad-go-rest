package utils

import (
	"context"
	"launchpad-go-rest/internal/lib/errors"
	"launchpad-go-rest/pkg/types/queue"
	"launchpad-go-rest/pkg/types/utils"
	"net/http"
)

func (s service) PublishTask(ctx context.Context, p utils.PublishTaskRequest) error {
	switch p.Queue {
	case queue.QueueTestQueue:
		return s.publisher.PublishTask(p.ID, p.Amount)
	case queue.QueueTestQueue2:
		return s.publisher.PublishTask2(p.ID, p.Amount)
	default:
		return errors.NewWithCode(http.StatusBadRequest, errors.INVALID_QUEUE_NAME, "Invalid queue name")
	}
}
