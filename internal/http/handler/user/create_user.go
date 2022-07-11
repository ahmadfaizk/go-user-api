package user

import (
	"net/http"
	"user-api/domain"
	"user-api/internal/helper"

	"github.com/labstack/echo/v4"
)

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
}

func (r *CreateUserRequest) toUser() *domain.User {
	return &domain.User{
		Name:     r.Name,
		Username: r.Username,
		Password: r.Password,
		Role:     domain.RoleUser,
	}
}

func CreateUserMethod(us domain.UserService) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		authUser := helper.ExtractUserFromContext(c)
		if authUser.Role != domain.RoleAdmin {
			return echo.ErrForbidden
		}

		var payload CreateUserRequest
		if err := c.Bind(&payload); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err := c.Validate(payload); err != nil {
			return err
		}

		user := payload.toUser()

		err := us.Create(ctx, user)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusCreated, user)
	}
}
