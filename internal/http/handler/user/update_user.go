package user

import (
	"net/http"
	"user-api/domain"
	"user-api/internal/helper"

	"github.com/labstack/echo/v4"
)

type UpdateUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
}

func (r *UpdateUserRequest) toUser() *domain.User {
	return &domain.User{
		Name:     r.Name,
		Username: r.Username,
		Password: r.Password,
	}
}

func UpdateUserMethod(us domain.UserService) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		id := c.Param("id")

		authUser := helper.ExtractUserFromContext(c)
		if authUser.Role != domain.RoleAdmin {
			return echo.ErrForbidden
		}

		var payload UpdateUserRequest
		if err := c.Bind(&payload); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err := c.Validate(payload); err != nil {
			return err
		}

		user := payload.toUser()

		err := us.Update(ctx, id, user)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, user)
	}
}
