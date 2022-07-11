package domain

import "context"

type AuthService interface {
	Login(ctx context.Context, username, password string) (*Token, error)
	RefreshTokenForUsername(ctx context.Context, username string) (*Token, error)
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
