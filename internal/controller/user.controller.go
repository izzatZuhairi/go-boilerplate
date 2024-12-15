package controller

import (
	"encoding/json"
	"net/http"
	"skeleton/common"
	"skeleton/internal/model"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := model.UserModel().GetAllUser()
	if err != nil {
		common.Json(w, http.StatusInternalServerError, err.Error(), false)
		return
	}

	common.Json(w, http.StatusOK, "List Users", users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var userData model.CreateUserData

	err := json.NewDecoder(r.Body).Decode(&userData)
	if err != nil {
		common.Json(w, http.StatusBadRequest, err.Error(), false)
		return
	}

	userId, err := model.UserModel().CreateUser(userData)
	if err != nil {
		common.Json(w, http.StatusInternalServerError, err.Error(), false)
	}

	common.Json(w, http.StatusOK, "User Id", userId)
}
