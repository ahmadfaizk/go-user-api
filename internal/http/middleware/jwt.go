package middleware

import (
	"user-api/config"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type JWTCustomClaims struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func WithJWTConfig(conf *config.Config) echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		Claims:        &JWTCustomClaims{},
		TokenLookup:   "header:" + echo.HeaderAuthorization,
		AuthScheme:    "Bearer",
		SigningKey:    []byte(conf.JWTAccessSecret),
		SigningMethod: middleware.AlgorithmHS256,
	}
	return middleware.JWTWithConfig(config)
}
