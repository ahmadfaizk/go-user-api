package auth

import (
	"fmt"
	"net/http"
	"user-api/config"
	"user-api/domain"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

func RefreshTokenMethod(conf *config.Config, as domain.AuthService) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		var payload RefreshTokenRequest

		if err := c.Bind(&payload); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		token, _ := jwt.Parse(payload.RefreshToken, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, echo.ErrInternalServerError
			}

			return []byte(conf.JWTRefreshSecret), nil
		})

		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			username := fmt.Sprintf("%v", claims["username"])
			newToken, err := as.RefreshTokenForUsername(ctx, username)
			if err != nil {
				return err
			}
			return c.JSON(http.StatusOK, newToken)
		}

		return echo.ErrInternalServerError
	}
}
