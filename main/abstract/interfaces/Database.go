package interfaces

import (
	"zhigalov_tutor_server_core/main/abstract/structs"
)

type Database interface {
	SelectUser(query *structs.User, args ...any) (*structs.User, error)
	SelectUsers(query *structs.User, args ...any) (*[]structs.User, error)
	CreateUser(user *structs.User) (*structs.User, error)
	UpdateUser(user *structs.User) (*structs.User, error)
}
