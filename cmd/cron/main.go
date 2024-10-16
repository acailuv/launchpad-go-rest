package main

import (
	"fmt"
	"launchpad-go-rest/internal/lib/config"
	"launchpad-go-rest/internal/lib/utils"
	"launchpad-go-rest/internal/repository"
	"launchpad-go-rest/internal/service"

	goerrors "github.com/go-errors/errors"
	"github.com/go-redis/redis"
	"github.com/hibiken/asynq"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"

	"github.com/jmoiron/sqlx"
	"github.com/robfig/cron"
)

func main() {
	config.Init()

	zerolog.ErrorStackMarshaler = func(err error) interface{} {
		frames := goerrors.Wrap(err, 1).StackFrames()

		stack := make([]string, len(frames))
		for i, frame := range frames {
			stack[i] = fmt.Sprintf("%s:%d", frame.File, frame.LineNumber)
		}

		return stack
	}

	db := sqlx.MustConnect("postgres", config.Configs.DatabaseDSN)
	defer db.Close()

	redis := redis.NewClient(&redis.Options{
		Addr: config.Configs.RedisDSN,
	})
	defer redis.Close()

	asynqClient := asynq.NewClient(asynq.RedisClientOpt{Addr: config.Configs.RedisDSN})
	defer asynqClient.Close()

	repositories := repository.Init(db, redis, asynqClient)
	utils := utils.New()
	_ = service.Init(repositories, utils)

	c := cron.New()
	c.AddFunc("* * * * * *", func() { fmt.Println("Heartbeat") })

	c.Run()
}
