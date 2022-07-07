package user

import (
	"net/http"
	"user-api/domain"

	"github.com/labstack/echo/v4"
)

func FetchUserMethod(us domain.UserService) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		data, err := us.Fetch(ctx)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, data)
	}
}
