package user

import (
	"net/http"
	"user-api/domain"

	"github.com/labstack/echo/v4"
)

type UpdateUserRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
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

		var payload UpdateUserRequest
		if err := c.Bind(&payload); err != nil {
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
