package userservice

import (
	"context"

	"github.com/andersoncarubelli/api-golang-postgresql-docker/internal/dto"
	"github.com/andersoncarubelli/api-golang-postgresql-docker/internal/handler/response"
	"github.com/andersoncarubelli/api-golang-postgresql-docker/internal/repository/userrepository"
)

func NewUserService(repository userrepository.UserRepository) UserService {
	return &service{
		repository,
	}
}

type service struct {
	repository userrepository.UserRepository
}

type UserService interface {
	CreateUser(ctx context.Context, u dto.CreateUserDto) error
	UpdateUser(ctx context.Context, u dto.UpdateUserDto, id string) error
	DeleteUser(ctx context.Context, id string) error
	GetUserByID(ctx context.Context, id string) (*response.UserResponse, error)
	FindManyUsers(ctx context.Context) (response.ManyUsersReponse, error)
	UpdateUserPassword(ctx context.Context, u *dto.UpdateUserPasswordDto, id string) error
}
