package auth

import (
	"context"
	"log"
	"time"
	"user-api/config"
	"user-api/domain"
	"user-api/internal/http/middleware"
	"user-api/tool"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type authService struct {
	config   *config.Config
	userRepo domain.UserRepository
}

func NewAuthService(config *config.Config, userRepo domain.UserRepository) domain.AuthService {
	return &authService{userRepo: userRepo, config: config}
}

func (as *authService) Login(ctx context.Context, username, password string) (*domain.Token, error) {
	user, err := as.userRepo.FindByUsername(ctx, username)
	if err != nil {
		return nil, echo.ErrUnauthorized
	}
	if !tool.CheckPasswordHash(password, user.Password) {
		return nil, echo.ErrUnauthorized
	}
	token, err := generateToken(as.config, *user)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (as *authService) RefreshTokenForUsername(ctx context.Context, username string) (*domain.Token, error) {
	user, err := as.userRepo.FindByUsername(ctx, username)
	if err != nil {
		return nil, echo.ErrNotFound
	}
	log.Println(user.Name)
	token, err := generateToken(as.config, *user)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func generateToken(config *config.Config, user domain.User) (*domain.Token, error) {
	claims := middleware.JWTCustomClaims{
		ID:       user.ID.Hex(),
		Name:     user.Name,
		Username: user.Username,
		Role:     user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	at, err := accessToken.SignedString([]byte(config.JWTAccessSecret))
	if err != nil {
		return nil, err
	}
	refreshClaims := middleware.JWTCustomClaims{
		ID:       user.ID.Hex(),
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(),
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	rt, err := refreshToken.SignedString([]byte(config.JWTRefreshSecret))
	if err != nil {
		return nil, err
	}
	token := domain.Token{
		AccessToken:  at,
		RefreshToken: rt,
	}
	return &token, nil
}
