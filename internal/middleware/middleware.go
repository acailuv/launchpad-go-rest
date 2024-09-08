package middleware

import "github.com/labstack/echo/v4"

type Middleware interface {
	VerifyToken() echo.MiddlewareFunc
}

type middleware struct{}

func Init() Middleware {
	return &middleware{}
}
