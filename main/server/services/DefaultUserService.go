package services

import (
	"zhigalov_tutor_server_core/main/abstract/interfaces/repos"
	"zhigalov_tutor_server_core/main/abstract/structs"
	"zhigalov_tutor_server_core/main/server/utils"
)

type DefaultUserService struct {
	repository *repos.UserRepository
}

func NewDefaultUserService(repository repos.UserRepository) *DefaultUserService {
	return &DefaultUserService{repository: &repository}
}

func (us *DefaultUserService) GetUsers(query *structs.User, args ...any) (*[]structs.User, error) {
	repo := *us.repository

	user, err := repo.ReadUsers(query, args...)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us *DefaultUserService) CreateUser(model *structs.UserRegisterModel) (*structs.User, error) {
	repo := *us.repository

	hashedPassword, err := utils.HashPassword(model.Password)
	if err != nil {
		return nil, err
	}
	model.Password = *hashedPassword

	user, err := repo.CreateUser(model)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us *DefaultUserService) UpdateUser(user *structs.User) (*structs.User, error) {
	//TODO implement me
	panic("implement me")
}

func (us *DefaultUserService) DeleteUser(query *any, args ...any) {
	//TODO implement me
	panic("implement me")
}
