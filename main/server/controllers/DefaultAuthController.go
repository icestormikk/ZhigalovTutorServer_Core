package controllers

import (
	"encoding/json"
	"net/http"
	"zhigalov_tutor_server_core/main/abstract/interfaces/services"
	"zhigalov_tutor_server_core/main/abstract/structs"
	"zhigalov_tutor_server_core/main/server/utils"
)

type DefaultAuthController struct {
	service *services.AuthService
}

func NewDefaultAuthController(service services.AuthService) *DefaultAuthController {
	return &DefaultAuthController{service: &service}
}

func (ac *DefaultAuthController) LoginUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	userLoginModel := &structs.UserLoginModel{}
	service := *ac.service

	err := decoder.Decode(userLoginModel)
	if err != nil {
		utils.SendResponse(w, "Bad Request", http.StatusBadRequest, &err)
		return
	}

	loggedInUser, err := service.LoginUser(userLoginModel)
	if err != nil {
		utils.SendResponse(w, err.Error(), http.StatusBadRequest, &err)
		return
	}

	utils.SendResponse(w, "Logged In", http.StatusOK, loggedInUser)
}

func (ac *DefaultAuthController) LogoutUser(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}
