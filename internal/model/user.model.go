package model

import (
	"context"
	"skeleton/common"
	database "skeleton/internal/db"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func UserModel() *BaseModel {
	mod := &BaseModel{
		ModelConstructor: &common.ModelConstructor{
			Collection: database.GetMongoDBConn("user").Collection("users"),
		},
	}

	return mod
}

type User struct {
	ID    bson.ObjectID `bson:"_id"`
	Name  string        `bson:"name"`
	Email string        `bson:"email"`
}

func (mod *BaseModel) GetAllUser() ([]User, error) {
	var results []User

	cur, err := mod.Collection.Find(context.TODO(), bson.D{})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		panic(err)
	}

	if err = cur.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	return results, nil
}
