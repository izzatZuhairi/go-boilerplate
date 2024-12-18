package model

import (
	"context"
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
	ID    bson.ObjectID `bson:"_id" json:"_id"`
	Name  string        `bson:"name" json:"name"`
	Email string        `bson:"email" json:"email"`
	User  bson.ObjectID `bson:"user" json:"user"`
}

type CreateStudentData struct {
	Name   string        `bson:"name" json:"name" validate:"required,max=100"`
	School string        `bson:"school" json:"school" validate:"required,max=100"`
	User   bson.ObjectID `bson:"user" json:"user" validate:"required,mongoid"`
}

func (mod *BaseModel) CreateStudent(body CreateStudentData) (interface{}, error) {
	inputMap := bson.M{
		"name":   body.Name,
		"school": body.School,
		"user":   body.User,
	}

	result, err := mod.Collection.InsertOne(context.TODO(), inputMap)
	if err != nil {
		return nil, err
	}

	return result.InsertedID, nil
}
