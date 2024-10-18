package queue_handler

import (
	"context"
	"encoding/json"
	"fmt"
	"launchpad-go-rest/pkg/types/queue"

	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

func (h handler) HandleTestQueue(ctx context.Context, t *asynq.Task) error {
	var p queue.TestQueuePayload

	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	log.Info().Msgf("Received Message: id=%s, amount=%s", p.ID, p.Amount)

	return nil
}
