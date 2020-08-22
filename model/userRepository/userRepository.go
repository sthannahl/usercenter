package userRepository

import (
	"context"
	"fmt"
	"sthannahl/usercenter/api/errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Name string
}

type UserRepository struct{}

var userRepository = &UserRepository{}

var collection *mongo.Collection

func GetInstance() *UserRepository {
	return userRepository
}

func (u *UserRepository) SetClient(c *mongo.Client) {
	collection = c.Database("user_center").Collection("user")
}

func (ur *UserRepository) FindOneUser() *User {
	var user User
	collection.FindOne(
		context.Background(),
		bson.D{}).Decode(&user)
	return &user
}

func (ur *UserRepository) Save(user *map[string]interface{}) error {
	fmt.Println(user)
	id, err := collection.InsertOne(
		context.Background(), user)
	if id == nil {
		err = errors.ErrExistUserName
	}
	return err
}
