package main

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestMongoRepo_CreateUser(t *testing.T) {
	ctx := context.Background()

	repo, err := NewMongoRepo(ctx, "mongodb://localhost:27017")
	if err != nil {
		t.Fatalf("error creating mongo repo: %s", err)
	}

	user := &User{
		ID:       primitive.NewObjectID(),
		Name:     "John",
		Email:    "john@example.com",
		Password: "password",
	}

	err = repo.CreateUser(ctx, user)
	if err != nil {
		t.Fatalf("error creating user: %s", err)
	}
}
