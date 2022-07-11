package auth

import (
	"net/http"
	"user-api/domain"

	"github.com/labstack/echo/v4"
)

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func LoginMethod(as domain.AuthService) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		var payload LoginRequest
		if err := c.Bind(&payload); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err := c.Validate(payload); err != nil {
			return err
		}

		token, err := as.Login(ctx, payload.Username, payload.Password)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, token)
	}
}
