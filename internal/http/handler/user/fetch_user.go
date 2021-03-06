package user

import (
	"net/http"
	"user-api/domain"
	"user-api/internal/helper"

	"github.com/labstack/echo/v4"
)

type FetchUserResponse struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func fetchUserResponse(u *domain.User) *FetchUserResponse {
	return &FetchUserResponse{
		ID:        u.ID.Hex(),
		Name:      u.Name,
		Username:  u.Username,
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func fetchUserResponses(us []*domain.User) []*FetchUserResponse {
	res := make([]*FetchUserResponse, len(us))
	for i, u := range us {
		res[i] = fetchUserResponse(u)
	}
	return res
}

func FetchUserMethod(us domain.UserService) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		authUser := helper.ExtractUserFromContext(c)
		if authUser.Role != domain.RoleAdmin {
			return echo.ErrForbidden
		}

		data, err := us.Fetch(ctx)
		if err != nil {
			return err
		}
		response := fetchUserResponses(data)

		return c.JSON(http.StatusOK, response)
	}
}
