package router

import (
	"user-api/config"
	"user-api/domain"
	"user-api/internal/http/handler/user"
	"user-api/internal/http/middleware"

	"github.com/labstack/echo/v4"
)

func NewUserRoute(echo *echo.Echo, config *config.Config, us domain.UserService) {
	jwtMiddleware := middleware.WithJWTConfig(config)
	route := echo.Group("/users")
	route.Use(jwtMiddleware)

	route.GET("/", user.FetchUserMethod(us))
	route.GET("/:id/", user.GetUserMethod(us))
	route.POST("/", user.CreateUserMethod(us))
	route.PUT("/:id/", user.UpdateUserMethod(us))
	route.DELETE("/:id/", user.DeleteUserMethod(us))
}
