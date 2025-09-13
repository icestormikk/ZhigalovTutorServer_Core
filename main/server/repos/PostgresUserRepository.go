package repos

import (
	"zhigalov_tutor_server_core/main/abstract/interfaces"
	"zhigalov_tutor_server_core/main/abstract/structs"
)

type PostgresUserRepository struct {
	database *interfaces.Database
}

func NewPostgresUserRepository(database interfaces.Database) *PostgresUserRepository {
	return &PostgresUserRepository{database: &database}
}

func (p *PostgresUserRepository) ReadUsers(query *structs.User, args ...any) (*[]structs.User, error) {
	users, err := (*p.database).SelectUsers(query, args)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (p *PostgresUserRepository) CreateUser(model *structs.UserRegisterModel) (*structs.User, error) {
	user := &structs.User{
		FirstName:      model.FirstName,
		LastName:       model.LastName,
		Email:          model.Email,
		HashedPassword: model.Password,
		BirthDate:      model.BirthDate,
	}

	db := *p.database
	createdUser, err := db.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (p *PostgresUserRepository) UpdateUser(user *structs.User) (*structs.User, error) {
	database := *p.database

	updatedUser, err := database.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (p *PostgresUserRepository) DeleteUser(query *structs.User, args ...any) {
	//TODO implement me
	panic("implement me")
}
