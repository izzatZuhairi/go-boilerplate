package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func GetMongoClient(uri string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		return nil, err
	}

	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	cancel()

	log.Println("Attempting to connect to DB")
	if err := client.Ping(context.Background(), nil); err != nil {
		panic(err)
	}

	return client, nil
}
