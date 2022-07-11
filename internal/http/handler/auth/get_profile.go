package auth

import (
	"net/http"
	"user-api/domain"
	"user-api/internal/helper"

	"github.com/labstack/echo/v4"
)

type GetProfileResponse struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func getProfileResponse(u *domain.User) *GetProfileResponse {
	return &GetProfileResponse{
		ID:        u.ID.Hex(),
		Name:      u.Name,
		Username:  u.Username,
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func GetProfileMethod(us domain.UserService) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		authUser := helper.ExtractUserFromContext(c)

		profile, err := us.FindByUsername(ctx, authUser.Username)
		if err != nil {
			return err
		}
		resp := getProfileResponse(profile)

		return c.JSON(http.StatusOK, resp)
	}
}
