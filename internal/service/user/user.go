package user

import (
	"context"
	"user-api/domain"
	"user-api/tool"
)

type userService struct {
	userRepo domain.UserRepository
}

func NewUserService(userRepo domain.UserRepository) domain.UserService {
	return &userService{userRepo: userRepo}
}

func (u *userService) Fetch(ctx context.Context) ([]*domain.User, error) {
	return u.userRepo.Fetch(ctx)
}

func (u *userService) FindById(ctx context.Context, id string) (*domain.User, error) {
	return u.userRepo.FindById(ctx, id)
}

func (u *userService) FindByUsername(ctx context.Context, username string) (*domain.User, error) {
	return u.userRepo.FindByUsername(ctx, username)
}

func (u *userService) Create(ctx context.Context, user *domain.User) error {
	hash, err := tool.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hash
	return u.userRepo.Create(ctx, user)
}

func (u *userService) Update(ctx context.Context, id string, user *domain.User) error {
	hash, err := tool.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hash
	return u.userRepo.Update(ctx, id, user)
}

func (u *userService) Delete(ctx context.Context, id string) error {
	return u.userRepo.Delete(ctx, id)
}
