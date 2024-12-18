package common

import "go.mongodb.org/mongo-driver/v2/mongo"

type ModelConstructor struct {
	Collection *mongo.Collection
}
