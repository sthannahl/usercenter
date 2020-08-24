package userrepository

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"sthannahl/usercenter/api/errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	buf    bytes.Buffer
	logger = log.New(&buf, "INFO: ", log.Lshortfile)

	infof = func(info string) {
		logger.Output(2, info)
	}
)

// UserRepository .
type UserRepository struct{}

var userRepository = &UserRepository{}

var collection *mongo.Collection

// GetInstance .
func GetInstance() *UserRepository {
	return userRepository
}

// SetClient .
func (u *UserRepository) SetClient(c *mongo.Client) {
	collection = c.Database("user_center").Collection("user")
}

// FindUserByTypeAndName .
func (u *UserRepository) FindUserByTypeAndName(typee, name string) *map[string]interface{} {
	var user map[string]interface{}
	collection.FindOne(
		context.Background(),
		bson.D{
			{"type", typee},
			{"user_id", name},
		}).Decode(&user)

	return &user
}

// Save .
func (u *UserRepository) Save(user *map[string]interface{}) error {
	id, err := collection.InsertOne(
		context.Background(), user)
	fmt.Println(id)
	fmt.Println(user)

	if err != nil {
		infof(err.Error())
	}
	if id == nil {
		err = errors.ErrExistUserName
	}
	return err
}
