package user

import (
	"net/http"
	"user-api/domain"
	"user-api/internal/helper"

	"github.com/labstack/echo/v4"
)

func DeleteUserMethod(us domain.UserService) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		id := c.Param("id")

		authUser := helper.ExtractUserFromContext(c)
		if authUser.Role != domain.RoleAdmin {
			return echo.ErrForbidden
		}

		err := us.Delete(ctx, id)
		if err != nil {
			return err
		}

		return c.NoContent(http.StatusOK)
	}
}
