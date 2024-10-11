package main

import (
	"fmt"
	"launchpad-go-rest/internal/controller"
	"launchpad-go-rest/internal/lib/config"
	_errors "launchpad-go-rest/internal/lib/errors"
	"launchpad-go-rest/internal/lib/utils"
	"launchpad-go-rest/internal/middleware"
	"launchpad-go-rest/internal/repository"
	"launchpad-go-rest/internal/router"
	"launchpad-go-rest/internal/service"
	"launchpad-go-rest/pkg/types/base"
	"net/http"

	_ "launchpad-go-rest/cmd/server/docs"

	goerrors "github.com/go-errors/errors"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	echo_middleware "github.com/labstack/echo/v4/middleware"
	echo_swagger "github.com/swaggo/echo-swagger"
)

// @title						Launchpad Go Rest API
// @version					1.0
// @description				This is a template for back end REST API server in Go.
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
// @host						localhost:1323
func main() {
	e := echo.New()
	config.Init()

	db := sqlx.MustConnect("postgres", config.Configs.DatabaseDSN)
	repositories := repository.Init(db)
	utils := utils.New()
	services := service.Init(repositories, utils)
	controllers := controller.Init(services)
	middleware := middleware.Init()
	router.Init(e, controllers, middleware)

	zerolog.ErrorStackMarshaler = func(err error) interface{} {
		frames := goerrors.Wrap(err, 1).StackFrames()

		stack := make([]string, len(frames))
		for i, frame := range frames {
			stack[i] = fmt.Sprintf("%s:%d", frame.File, frame.LineNumber)
		}

		return stack
	}
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	e.Use(echo_middleware.RequestLoggerWithConfig(echo_middleware.RequestLoggerConfig{
		LogMethod: true,
		LogURI:    true,
		LogValuesFunc: func(c echo.Context, v echo_middleware.RequestLoggerValues) error {
			log.Info().
				Str("URI", fmt.Sprintf("%s %s", v.Method, v.URI)).
				Msg("request incoming")

			return nil
		},
	}))
	e.Use(echo_middleware.Recover())

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		resp := base.Response[any]{
			Error: &base.ErrorResponse{
				Code:    _errors.INTERNAL_SERVER_ERROR,
				Message: "Internal Server Error",
			},
		}

		var statusCode int
		if e, ok := err.(*goerrors.Error); ok {
			if ex, ok := e.Unwrap().(*_errors.Error); ok {
				statusCode = ex.StatusCode
				resp.Error.Code = ex.ErrorCode

				if statusCode != http.StatusInternalServerError {
					resp.Error.Message = e.Err.Error()
				}
			}
		} else if e, ok := err.(*echo.HTTPError); ok {
			statusCode = e.Code
			resp.Error.Code = e.Code
			resp.Error.Message = e.Message.(string)
		}

		log.Error().Stack().Err(err).Send()
		c.JSON(statusCode, resp)
	}

	// Swagger
	e.GET("/swagger*", echo_swagger.WrapHandler)

	e.HideBanner = true
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", config.Configs.Port)))
}
