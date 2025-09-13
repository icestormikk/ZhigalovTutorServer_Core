package repos

import "zhigalov_tutor_server_core/main/abstract/structs"

type UserRepository interface {
	CreateUser(user *structs.UserRegisterModel) (*structs.User, error)
	ReadUsers(query *structs.User, args ...any) (*[]structs.User, error)
	UpdateUser(user *structs.User) (*structs.User, error)
	DeleteUser(query *structs.User, args ...any)
}
