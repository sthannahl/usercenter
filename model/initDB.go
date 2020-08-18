package model

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Mongo *mongo.Client
}

var DB *Database

//初始化
func InitDB(uri string) {
	DB = &Database{
		Mongo: SetConnect(uri),
	}
}

// 连接设置
func SetConnect(uri string) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri).SetMaxPoolSize(20)) // 连接池
	if err != nil {
		fmt.Println(err)
	}
	return client
}
