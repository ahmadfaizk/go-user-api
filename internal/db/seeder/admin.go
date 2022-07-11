package seeder

import (
	"context"
	"user-api/domain"
)

func RunAdminSeeder(userService domain.UserService) error {
	name := "Admin"
	username := "admin"
	password := "admin123"
	user, _ := userService.FindByUsername(context.TODO(), username)
	if user == nil {
		user = &domain.User{
			Name:     name,
			Username: username,
			Password: password,
			Role:     domain.RoleAdmin,
		}
		err := userService.Create(context.TODO(), user)
		return err
	}
	return nil
}
