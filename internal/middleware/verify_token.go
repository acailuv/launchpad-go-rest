package middleware

import (
	"launchpad-go-rest/internal/lib/config"
	"launchpad-go-rest/internal/lib/errors"
	"net/http"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func (m middleware) VerifyToken() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		ErrorHandler: func(c echo.Context, err error) error {
			return errors.NewWithCode(http.StatusUnauthorized, errors.INVALID_AUTH_TOKEN, "Unauthorized")
		},
		SigningKey: []byte(config.Configs.JWTSecret),
	})
}
