package publisher

import (
	"github.com/acailuv/numeric"
	"github.com/hibiken/asynq"
)

type Repository interface {
	PublishTask(id string, amount numeric.Numeric) error
	PublishTask2(id string, amount numeric.Numeric) error
}

type repository struct {
	client *asynq.Client
}

func New(client *asynq.Client) Repository {
	return &repository{
		client: client,
	}
}
