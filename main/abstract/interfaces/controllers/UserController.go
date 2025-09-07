package controllers

import (
	"net/http"
	"zhigalov_tutor_server_core/main/abstract/structs"
)

type UserController interface {
	GetUser(w http.ResponseWriter, r *http.Request) (*structs.User, error)
	CreateUser(w http.ResponseWriter, r *http.Request) (*structs.User, error)
	UpdateUser(w http.ResponseWriter, r *http.Request) (*structs.User, error)
	DeleteUser(w http.ResponseWriter, r *http.Request)
}
