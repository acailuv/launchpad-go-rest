package publisher

import (
	"encoding/json"
	"launchpad-go-rest/internal/lib/errors"
	"launchpad-go-rest/pkg/types/queue"

	"github.com/acailuv/numeric"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

func (r repository) PublishTask2(id string, amount numeric.Numeric) error {
	payload, err := json.Marshal(queue.TestQueuePayload{ID: id, Amount: amount})
	if err != nil {
		return errors.New(err)
	}

	task := asynq.NewTask(queue.QueueTestQueue2, payload, asynq.Queue(queue.QueueTestQueue2))
	info, err := r.client.Enqueue(task)
	if err != nil {
		return errors.New(err)
	}

	log.Info().Any("queue", info.Queue).Str("payload", string(info.Payload)).Msg("Published successfully!")

	return nil
}
