package db

import (
	"context"
	"fmt"
	"user-api/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewDatabase(conf *config.Config) (*mongo.Database, error) {
	clientOptions := options.Client()
	credentials := options.Credential{
		Username: conf.DatabaseUser,
		Password: conf.DatabasePassword,
	}
	uri := fmt.Sprintf("mongodb://%s:%d", conf.DatabaseHost, conf.DatabasePort)
	clientOptions.ApplyURI(uri).SetAuth(credentials)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}
	err = client.Connect(context.TODO())
	if err != nil {
		return nil, err
	}
	return client.Database(conf.DatabaseName), nil
}
