package model

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Name string
}

type UserRepository struct{}

var client *mongo.Client

func (u *UserRepository) SetClient(c *mongo.Client) {
	client = c
}

func (ur *UserRepository) FindOneUser() *User {
	var user User
	collection := client.Database("user_center").Collection("user")
	collection.FindOne(
		context.Background(),
		bson.D{}).Decode(&user)
	return &user
}
