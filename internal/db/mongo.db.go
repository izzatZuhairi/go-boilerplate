package database

import (
	"context"
	"skeleton/config"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
	"go.uber.org/zap"
)

var connections = make(map[string]*mongo.Database)

func connectMongoDB(connStr string, dbName string, logger *zap.SugaredLogger) (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI(connStr)
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	cancel()

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	db := client.Database(dbName)

	logger.Infow("Connected to db successfully", "db", dbName)

	return db, nil
}

func initMongoDB(dbName string, connStr string, logger *zap.SugaredLogger) {
	if _, exists := connections[dbName]; exists {
		logger.Infow("Mongodb connection already exists, skipping initialization", "db", dbName)
		return
	}

	db, err := connectMongoDB(connStr, dbName, logger)
	if err != nil {
		logger.Fatalf("No mongodb connection found for database '%s'", dbName)
	}

	connections[dbName] = db
}

func InitDBConnections(logger *zap.SugaredLogger) {
	initMongoDB("user", config.LoadConfig("MONGO_USER_CONN_STR"), logger)
	initMongoDB("student", config.LoadConfig("MONGO_STUDENT_CONN_STR"), logger)
}

func GetMongoDBConn(dbName string) *mongo.Database {
	if db, exists := connections[dbName]; exists {
		return db
	}

	return nil
}
