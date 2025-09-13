package controllers

import (
	"encoding/json"
	"net/http"
	"zhigalov_tutor_server_core/main/abstract/interfaces/services"
	"zhigalov_tutor_server_core/main/abstract/structs"
	"zhigalov_tutor_server_core/main/server/utils"
)

type DefaultUserController struct {
	service *services.UserService
}

func NewDefaultUserController(service services.UserService) *DefaultUserController {
	return &DefaultUserController{service: &service}
}

func (uc *DefaultUserController) GetUser(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (uc *DefaultUserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	service := *uc.service

	users, err := service.GetUsers(nil)
	if err != nil {
		utils.SendResponse[error](w, "Error occurred", http.StatusInternalServerError, &err)
		return
	}

	utils.SendResponse[[]structs.User](w, "Success", http.StatusOK, users)
}

func (uc *DefaultUserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	service := *uc.service
	decoder := json.NewDecoder(r.Body)

	model := &structs.UserRegisterModel{}
	err := decoder.Decode(model)
	if err != nil {
		utils.SendResponse[error](w, "Body not suitable", http.StatusBadRequest, &err)
		return
	}

	user, err := service.CreateUser(model)
	if err != nil {
		utils.SendResponse[error](w, "Error occurred", http.StatusInternalServerError, &err)
		return
	}

	utils.SendResponse[structs.User](w, "Created", http.StatusCreated, user)
}

func (uc *DefaultUserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (uc *DefaultUserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}
