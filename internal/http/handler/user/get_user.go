package user

import (
	"net/http"
	"user-api/domain"

	"github.com/labstack/echo/v4"
)

func GetUserMethod(us domain.UserService) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		id := c.Param("id")

		data, err := us.FindById(ctx, id)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, data)
	}
}
