package controller

import (
	"net/http"
	"skeleton/common"
	"skeleton/internal/model"
	"skeleton/internal/service"
	"skeleton/internal/validator"
	"skeleton/types"
)

func (c *Controller) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := service.GetAllUser()
	if err != nil {
		common.Json(w, http.StatusInternalServerError, err.Error(), false)
		return
	}

	common.Json(w, http.StatusOK, "List Users", users)
}

func (c *Controller) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userData model.CreateUserData

	if err := validator.ReadJSON(w, r, &userData); err != nil {
		common.Json(w, http.StatusBadRequest, err.Error(), false)
		return
	}

	if err := validator.Validate.Struct(&userData); err != nil {
		common.Json(w, http.StatusBadRequest, err.Error(), false)
		return
	}

	userId, err := service.CreateUser(userData)
	if err != nil {
		common.Json(w, http.StatusInternalServerError, err.Error(), false)
		return
	}

	common.Json(w, http.StatusOK, "User Id", userId)
}

func (c *Controller) CreateStudentAndUser(w http.ResponseWriter, r *http.Request) {
	var input types.CreateUserAndStudent

	if err := validator.ReadJSON(w, r, &input); err != nil {
		common.Json(w, http.StatusBadRequest, err.Error(), false)
		return
	}

	if err := validator.Validate.Struct(&input); err != nil {
		common.Json(w, http.StatusBadRequest, err.Error(), false)
		return
	}

	result, err := service.CreateUserAndStudent(input)
	if err != nil {
		common.Json(w, http.StatusInternalServerError, err.Error(), false)
		return
	}

	common.Json(w, http.StatusOK, "User Id", result)
}
