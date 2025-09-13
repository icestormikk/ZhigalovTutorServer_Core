package services

import "zhigalov_tutor_server_core/main/abstract/structs"

type AuthService interface {
	LoginUser(*structs.UserLoginModel) (*structs.User, error)
	LogoutUser(*structs.UserLoginModel) error
}
