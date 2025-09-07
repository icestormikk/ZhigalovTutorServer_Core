package interfaces

import "zhigalov_tutor_server_core/main/abstract/structs"

type Database interface {
	SelectUser(query any, args ...any) (*structs.User, error)
	SelectUsers(query *any, args ...any) (*[]structs.User, error)
}
