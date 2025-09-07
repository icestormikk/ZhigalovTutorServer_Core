package services

import (
	"zhigalov_tutor_server_core/main/abstract/interfaces/repos"
	"zhigalov_tutor_server_core/main/abstract/structs"
)

type DefaultUserService struct {
	repository *repos.UserRepository
}

func (us *DefaultUserService) SelectUsers(query any, args ...any) (*[]structs.User, error) {
	//TODO implement me
	panic("implement me")
}

func NewDefaultUserService(repository repos.UserRepository) *DefaultUserService {
	return &DefaultUserService{repository: &repository}
}

func (us *DefaultUserService) GetUsers(query *any, args ...any) (*[]structs.User, error) {
	repo := *us.repository

	user, err := repo.ReadUsers(query, args...)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us *DefaultUserService) CreateUser(user structs.User) (*structs.User, error) {
	//TODO implement me
	panic("implement me")
}

func (us *DefaultUserService) UpdateUser(user structs.User) (*structs.User, error) {
	//TODO implement me
	panic("implement me")
}

func (us *DefaultUserService) DeleteUser(query *any, args ...any) {
	//TODO implement me
	panic("implement me")
}
