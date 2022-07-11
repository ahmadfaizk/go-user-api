package helper

import (
	"user-api/internal/http/middleware"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func ExtractUserFromContext(c echo.Context) *middleware.JWTCustomClaims {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*middleware.JWTCustomClaims)
	return claims
}
