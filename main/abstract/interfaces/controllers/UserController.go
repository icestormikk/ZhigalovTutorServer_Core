package controllers

import (
	"net/http"
)

type UserController interface {
	GetUser(http.ResponseWriter, *http.Request)
	GetUsers(http.ResponseWriter, *http.Request)
	CreateUser(http.ResponseWriter, *http.Request)
	UpdateUser(http.ResponseWriter, *http.Request)
	DeleteUser(http.ResponseWriter, *http.Request)
}
