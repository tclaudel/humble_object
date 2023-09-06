package main

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ErrConnectingToMongoDatabase = errors.New("error connecting to mongo database")
	ErrInsertingUser             = errors.New("error inserting user")
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name,omitempty"`
	Email    string             `bson:"email,omitempty"`
	Password string             `bson:"password,omitempty"`
}

type MongoRepo struct {
	collection *mongo.Collection
}

func NewMongoRepo(ctx context.Context, mongoURI string) (*MongoRepo, error) {
	const (
		dbName         = "test"
		collectionName = "users"
	)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrConnectingToMongoDatabase, err)
	}

	collection := client.Database(dbName).Collection(collectionName)

	return &MongoRepo{
		collection: collection,
	}, nil
}

func (m *MongoRepo) CreateUser(ctx context.Context, user *User) error {
	_, err := m.collection.InsertOne(ctx, user)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrInsertingUser, err)
	}

	return nil
}
