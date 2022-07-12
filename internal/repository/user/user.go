package user

import (
	"context"
	"errors"
	"log"
	"time"
	"user-api/domain"
	"user-api/pkg/constant"

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

func (u *userRepository) Fetch(ctx context.Context) ([]*domain.User, error) {
	csr, err := u.db.Collection("users").Find(ctx, bson.M{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer csr.Close(ctx)

	result := make([]*domain.User, 0)
	for csr.Next(ctx) {
		var row domain.User
		err := csr.Decode(&row)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		result = append(result, &row)
	}
	return result, nil
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

func (u *userRepository) FindByUsername(ctx context.Context, username string) (*domain.User, error) {
	selector := bson.M{"username": username}
	result := u.db.Collection("users").FindOne(ctx, selector)
	var user domain.User
	err := result.Decode(&user)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &user, nil
}

func (u *userRepository) Create(ctx context.Context, user *domain.User) error {
	userExist, _ := u.FindByUsername(ctx, user.Username)
	if userExist != nil {
		return errors.New("username is already exist")
	}

	t := time.Now().Format(constant.LayoutDateTimeISO8601)
	user.ID = primitive.NewObjectID()
	user.CreatedAt = t
	user.UpdatedAt = t
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
	t := time.Now().Format(constant.LayoutDateTimeISO8601)
	user.ID = objectId
	user.UpdatedAt = t
	selector := bson.M{"_id": objectId}
	userData := bson.M{
		"name":       user.Name,
		"username":   user.Username,
		"password":   user.Password,
		"updated_at": user.UpdatedAt,
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
