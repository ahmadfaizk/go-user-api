package user

import (
	"context"
	"log"
	"user-api/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	db *mongo.Database
}

func NewUserRepository(db *mongo.Database) domain.UserRepository {
	return &userRepository{db: db}
}

func (u *userRepository) Fetch(ctx context.Context) (*[]domain.User, error) {
	csr, err := u.db.Collection("users").Find(ctx, bson.M{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer csr.Close(ctx)

	result := make([]domain.User, 0)
	for csr.Next(ctx) {
		var row domain.User
		err := csr.Decode(&row)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		result = append(result, row)
	}
	return &result, nil
}

func (u *userRepository) FindById(ctx context.Context, id string) (*domain.User, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	selector := bson.M{"_id": objectId}
	result := u.db.Collection("users").FindOne(ctx, selector)
	var user domain.User
	err = result.Decode(&user)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &user, nil
}

func (u *userRepository) Create(ctx context.Context, user *domain.User) error {
	_, err := u.db.Collection("users").InsertOne(ctx, user)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (u *userRepository) Update(ctx context.Context, id string, user *domain.User) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return err
	}
	user.ID = objectId
	selector := bson.M{"_id": objectId}
	userData := bson.M{
		"name":     user.Name,
		"username": user.Password,
		"password": user.Password,
	}
	_, err = u.db.Collection("users").UpdateOne(ctx, selector, bson.M{"$set": userData})
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (u *userRepository) Delete(ctx context.Context, id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return err
	}
	selector := bson.M{"_id": objectId}
	_, err = u.db.Collection("users").DeleteOne(ctx, selector)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
