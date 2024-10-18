package main

import (
	"launchpad-go-rest/internal/lib/config"
	"launchpad-go-rest/internal/lib/utils"
	"launchpad-go-rest/internal/queue_handler"
	"launchpad-go-rest/internal/repository"
	"launchpad-go-rest/internal/service"
	"launchpad-go-rest/pkg/types/queue"

	"github.com/go-redis/redis"
	"github.com/hibiken/asynq"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

func main() {
	config.Init()

	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: config.Configs.RedisDSN},
		asynq.Config{
			Concurrency: config.Configs.QueuesConcurrency,
			Queues:      queue.AllQueues,
		},
	)

	db := sqlx.MustConnect("postgres", config.Configs.DatabaseDSN)
	defer db.Close()

	redis := redis.NewClient(&redis.Options{
		Addr: config.Configs.RedisDSN,
	})
	defer redis.Close()

	client := asynq.NewClient(asynq.RedisClientOpt{Addr: config.Configs.RedisDSN})
	defer client.Close()

	repositories := repository.Init(db, redis, client)
	utils := utils.New()
	services := service.Init(repositories, utils)

	handlers := queue_handler.New(repositories, services)

	mux := asynq.NewServeMux()
	mux.HandleFunc(queue.QueueTestQueue, handlers.HandleTestQueue)
	mux.HandleFunc(queue.QueueTestQueue2, handlers.HandleTestQueue2)

	if err := srv.Run(mux); err != nil {
		log.Error().Err(err).Msg("could not run server")
	}
}
