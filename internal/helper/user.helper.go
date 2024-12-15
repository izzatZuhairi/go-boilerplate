package helper

// import (
// 	"context"
// 	model "skeleton/internal/model"
//
// 	"go.mongodb.org/mongo-driver/v2/bson"
// 	"go.mongodb.org/mongo-driver/v2/mongo"
// )
//
// func (usermod *model.BaseModel) GetAllUsers() ([]model.User, error) {
// 	var results []model.User
//
// 	cur, err := usermod.Collection.Find(context.TODO(), bson.D{})
// 	if err != nil {
// 		if err == mongo.ErrNoDocuments {
// 			return
// 		}
// 		panic(err)
// 	}
// }
