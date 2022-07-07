package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name     string             `bson:"name" json:"name"`
	Username string             `bson:"username" json:"username"`
	Password string             `bson:"password" json:"password"`
	UserType string             `bson:"user_type" json:"user_type"`
}

type UserService interface {
	Fetch(ctx context.Context) (*[]User, error)
	FindById(ctx context.Context, id string) (*User, error)
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, id string, user *User) error
	Delete(ctx context.Context, id string) error
}

type UserRepository interface {
	Fetch(ctx context.Context) (*[]User, error)
	FindById(ctx context.Context, id string) (*User, error)
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, id string, user *User) error
	Delete(ctx context.Context, id string) error
}
