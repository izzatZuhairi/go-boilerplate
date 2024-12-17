package service

import (
	"skeleton/internal/model"
	"skeleton/types"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func GetAllUser() ([]model.User, error) {
	users, err := model.UserModel().GetAllUser()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func CreateUser(input model.CreateUserData) (interface{}, error) {
	userId, err := model.UserModel().CreateUser(input)
	if err != nil {
		return nil, err
	}

	return userId, nil
}

func CreateUserAndStudent(input types.CreateUserAndStudent) (types.CreateUserAndStudentRes, error) {
	createUserData := model.CreateUserData{
		Name:  input.Name,
		Email: input.Email,
	}

	var response types.CreateUserAndStudentRes

	userId, err := model.UserModel().CreateUser(createUserData)
	if err != nil {
		return response, err
	}

	createStudentData := model.CreateStudentData{
		Name:   input.Name,
		School: input.School,
		User:   userId.(bson.ObjectID),
	}

	studentId, err := model.StudentModel().CreateStudent(createStudentData)
	if err != nil {
		return response, err
	}

	response = types.CreateUserAndStudentRes{
		UserId:    userId.(bson.ObjectID).Hex(),
		StudentId: studentId.(bson.ObjectID).Hex(),
	}

	return response, nil
}
