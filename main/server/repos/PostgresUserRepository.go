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

func (p *PostgresUserRepository) ReadUsers(query *any, args ...any) (*[]structs.User, error) {
	users, err := (*p.database).SelectUsers(query, args)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (p *PostgresUserRepository) CreateUser(user structs.User) (*structs.User, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PostgresUserRepository) UpdateUser(user structs.User) (*structs.User, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PostgresUserRepository) DeleteUser(query *any, args ...any) {
	//TODO implement me
	panic("implement me")
}
