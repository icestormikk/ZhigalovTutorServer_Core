package services

import "zhigalov_tutor_server_core/main/abstract/structs"

type UserService interface {
	GetUsers(query *structs.User, args ...any) (*[]structs.User, error)
	CreateUser(user *structs.UserRegisterModel) (*structs.User, error)
	UpdateUser(user *structs.User) (*structs.User, error)
	DeleteUser(query *any, args ...any)
}
