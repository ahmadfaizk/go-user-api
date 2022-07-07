package user

import (
	"net/http"
	"user-api/domain"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateUserRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *CreateUserRequest) toUser() *domain.User {
	return &domain.User{
		ID:       primitive.NewObjectID(),
		Name:     r.Name,
		Username: r.Username,
		Password: r.Password,
	}
}

func CreateUserMethod(us domain.UserService) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		var payload CreateUserRequest
		if err := c.Bind(&payload); err != nil {
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
