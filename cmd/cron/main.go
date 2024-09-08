package main

import (
	"fmt"
	"launchpad-go-rest/internal/lib/config"
	"launchpad-go-rest/internal/repository"
	"launchpad-go-rest/internal/service"

	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/robfig/cron"
)

func main() {
	e := echo.New()
	config.Init()

	db := sqlx.MustConnect("postgres", config.Configs.DatabaseDSN)
	repositories := repository.Init(db, e.Logger)
	_ = service.Init(repositories, e.Logger)

	c := cron.New()
	c.AddFunc("* * * * * *", func() { fmt.Println("Heartbeat") })

	c.Run()
}
