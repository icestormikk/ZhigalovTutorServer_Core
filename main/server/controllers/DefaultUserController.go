package controllers

import (
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

func (uc *DefaultUserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	service := *uc.service

	users, err := service.GetUsers(nil)
	if err != nil {
		utils.SendResponse[error](w, "Error occurred", http.StatusInternalServerError, &err)
		return
	}

	utils.SendResponse[[]structs.User](w, "Success", http.StatusOK, users)
}

func (uc *DefaultUserController) CreateUser(w http.ResponseWriter, r *http.Request) (*structs.User, error) {
	//TODO implement me
	panic("implement me")
}

func (uc *DefaultUserController) UpdateUser(w http.ResponseWriter, r *http.Request) (*structs.User, error) {
	//TODO implement me
	panic("implement me")
}

func (uc *DefaultUserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}
