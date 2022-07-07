package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	e.Pre(middleware.AddTrailingSlash())
	e.Use(middleware.CORS())

	return e
}
