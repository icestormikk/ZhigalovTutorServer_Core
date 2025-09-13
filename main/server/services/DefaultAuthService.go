package services

import (
	"time"
	"zhigalov_tutor_server_core/main/abstract/interfaces/repos"
	"zhigalov_tutor_server_core/main/abstract/structs"
)

type DefaultAuthService struct {
	repo *repos.UserRepository
}

func NewDefaultAuthService(repo repos.UserRepository) *DefaultAuthService {
	return &DefaultAuthService{repo: &repo}
}

func (as *DefaultAuthService) LoginUser(model *structs.UserLoginModel) (*structs.User, error) {
	repo := *as.repo
	users, err := repo.ReadUsers(&structs.User{Email: model.Email}, nil)
	if err != nil {
		return nil, err
	}

	user := (*users)[0]

	user.LastLoginDate = time.Now()
	loggedUser, err := repo.UpdateUser(&user)
	if err != nil {
		return nil, err
	}

	return loggedUser, err
}

func (as *DefaultAuthService) LogoutUser(model *structs.UserLoginModel) error {
	//TODO implement me
	panic("implement me")
}
