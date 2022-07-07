package user

import (
	"net/http"
	"user-api/domain"

	"github.com/labstack/echo/v4"
)

func DeleteUserMethod(us domain.UserService) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		id := c.Param("id")

		err := us.Delete(ctx, id)
		if err != nil {
			return err
		}

		return c.NoContent(http.StatusOK)
	}
}
