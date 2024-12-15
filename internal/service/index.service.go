package service

import "skeleton/internal/model"

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
