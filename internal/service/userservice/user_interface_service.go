package userservice

import "github.com/andersoncarubelli/api-golang-postgresql-docker/internal/repository/userrepository"

func NewUserService(repository userrepository.UserRepository) UserService {
	return &service{
		repository,
	}
}

type service struct {
	repository userrepository.UserRepository
}

type UserService interface {
	CreateUser() error
}
