package model

import (
	"skeleton/common"
	database "skeleton/internal/db"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func StudentModel() *BaseModel {
	mod := &BaseModel{
		ModelConstructor: &common.ModelConstructor{
			Collection: database.GetMongoDBConn("student").Collection("students"),
		},
	}

	return mod
}

type Student struct {
	ID    bson.ObjectID `bson:"_id"`
	Name  string        `bson:"name"`
	Email string        `bson:"email"`
	User  bson.ObjectID `bson:"user"`
}
