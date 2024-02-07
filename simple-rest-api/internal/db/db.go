package db

import (
	"context"
	"os"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoInstance *mongo.Client
var mongoError error

var syncOnce sync.Once

const (
	Db = "mydb"
)

func GetClient() (*mongo.Client, error) {
	syncOnce.Do(func() {
		dbUrl := os.Getenv("DB_URL")
		options := options.Client().ApplyURI(dbUrl)

		client, err := mongo.Connect(context.TODO(), options)

		if err != nil {
			mongoError = err
		}

		mongoInstance = client
	})

	return mongoInstance, mongoError
}
