package router

import (
	"user-api/config"
	"user-api/domain"
	"user-api/internal/http/handler/auth"
	"user-api/internal/http/middleware"

	"github.com/labstack/echo/v4"
)

func NewAuthRoute(echo *echo.Echo, config *config.Config, as domain.AuthService, us domain.UserService) {
	route := echo.Group("/auth")
	jwtMiddleware := middleware.WithJWTConfig(config)

	route.POST("/login/", auth.LoginMethod(as))
	route.POST("/refresh/", auth.RefreshTokenMethod(config, as))
	route.GET("/profile/", auth.GetProfileMethod(us), jwtMiddleware)
}
