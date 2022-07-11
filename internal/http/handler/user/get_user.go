package user

import (
	"net/http"
	"user-api/domain"
	"user-api/internal/helper"

	"github.com/labstack/echo/v4"
)

func GetUserMethod(us domain.UserService) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		id := c.Param("id")

		authUser := helper.ExtractUserFromContext(c)
		if authUser.Role == domain.RoleUser {
			if authUser.ID != id {
				return echo.ErrForbidden
			}
		}

		data, err := us.FindById(ctx, id)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound)
		}

		response := fetchUserResponse(data)

		return c.JSON(http.StatusOK, response)
	}
}
