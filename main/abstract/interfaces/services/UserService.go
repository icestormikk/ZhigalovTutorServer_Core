package services

import "zhigalov_tutor_server_core/main/abstract/structs"

type UserService interface {
	GetUsers(query *any, args ...any) (*[]structs.User, error)
	CreateUser(user structs.User) (*structs.User, error)
	UpdateUser(user structs.User) (*structs.User, error)
	DeleteUser(query *any, args ...any)
}
