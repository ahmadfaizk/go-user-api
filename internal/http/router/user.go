package router

import (
	"user-api/domain"
	"user-api/internal/http/handler/user"

	"github.com/labstack/echo/v4"
)

func NewUserRoute(echo *echo.Echo, us domain.UserService) {
	route := echo.Group("/users")
	route.GET("/", user.FetchUserMethod(us))
	route.GET("/:id/", user.GetUserMethod(us))
	route.POST("/", user.CreateUserMethod(us))
	route.PUT("/:id/", user.UpdateUserMethod(us))
	route.DELETE("/:id/", user.DeleteUserMethod(us))
}
