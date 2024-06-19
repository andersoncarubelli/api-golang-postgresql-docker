package userservice

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/andersoncarubelli/api-golang-postgresql-docker/internal/dto"
	"github.com/andersoncarubelli/api-golang-postgresql-docker/internal/entity"
	"github.com/andersoncarubelli/api-golang-postgresql-docker/internal/handler/response"
	"github.com/google/uuid"
	"github.com/wiliamvj/api-users-golang/api/viacep"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) CreateUser(ctx context.Context, u dto.CreateUserDto) error {
	userExists, err := s.repository.FindUserByEmail(ctx, u.Email)

	if err != nil {
		slog.Error("error to search user by email", "err", err, slog.String("package", "userservice"))
		return err
	}

	if userExists != nil {
		slog.Error("user already exists", slog.String("package", "userservice"))
		return errors.New("user already exists")
	}

	passwordEncrypted, err := bcrypt.GenerateFromPassword([]byte(u.Password), 12)
	if err != nil {
		slog.Error("error to encrypt password", "err", err, slog.String("package", "userservice"))
		return errors.New("error to encrypt password")
	}

	cep, err := viacep.GetCep(u.CEP)
	if err != nil {
		slog.Error("error to get cep", "err", err, slog.String("package", "userservice"))
		return err
	}

	newUser := entity.UserEntity{
		ID:       uuid.New().String(),
		Name:     u.Name,
		Email:    u.Email,
		Password: string(passwordEncrypted),
		Address: entity.UserAddress{
			CEP:        cep.CEP,
			IBGE:       cep.IBGE,
			UF:         cep.UF,
			City:       cep.City,
			Complement: cep.Complement,
			Street:     cep.Street,
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = s.repository.CreateUser(ctx, &newUser)
	if err != nil {
		slog.Error("error to create user", "err", err, slog.String("package", "userservice"))
		return err
	}
	return nil
}

func (s *service) UpdateUser(ctx context.Context, u dto.UpdateUserDto, id string) error {
	userExists, err := s.repository.FindUserByID(ctx, id)
	if err != nil {
		slog.Error("error to search user by id", "err", err, slog.String("package", "userservice"))
		return err
	}

	if userExists == nil {
		slog.Error("user not found", slog.String("package", "userservice"))
		return errors.New("user not exists")
	}

	var updateUser entity.UserEntity

	if u.Email != "" {
		verifyUserEmail, err := s.repository.FindUserByEmail(ctx, u.Email)
		if err != nil {
			slog.Error("error ro search user by email", "err", err, slog.String("package", "userservice"))
			return err
		}

		if verifyUserEmail != nil {
			slog.Error("email already in use", slog.String("package", "userservice"))
			return errors.New("email already in use")
		}

		updateUser.Email = u.Email
	}

	if u.CEP != "" {
		cep, err := viacep.GetCep(u.CEP)
		if err != nil {
			slog.Error("error to get cep", "err", err, slog.String("package", "userservice"))
			return err
		}

		updateUser.Address = entity.UserAddress{
			CEP:        cep.CEP,
			IBGE:       cep.IBGE,
			UF:         cep.UF,
			City:       cep.City,
			Complement: cep.Complement,
			Street:     cep.Street,
		}
	}

	updateUser.ID = id
	updateUser.Name = u.Name
	updateUser.UpdatedAt = time.Now()

	err = s.repository.UpdateUser(ctx, &updateUser)
	if err != nil {
		slog.Error("error to update user", "err", err, slog.String("package", "userservice"))
		return err
	}

	return nil
}

func (s *service) DeleteUser(ctx context.Context, id string) error {
	userExists, err := s.repository.FindUserByID(ctx, id)
	if err != nil {
		slog.Error("error to search user by id", "err", err, slog.String("package", "userservice"))
		return err
	}

	if userExists == nil {
		slog.Error("user not found", slog.String("package", "userservice"))
		return errors.New("user not found")
	}

	err = s.repository.DeleteUser(ctx, id)
	if err != nil {
		slog.Error("error to delete user", "err", err, slog.String("package", "userservice"))
		return err
	}

	return nil
}

func (s *service) GetUserByID(ctx context.Context, id string) (*response.UserResponse, error) {
	userExists, err := s.repository.FindUserByID(ctx, id)
	if err != nil {
		slog.Error("error to search user by id", "err", err, slog.String("package", "userservice"))
		return nil, err
	}

	if userExists == nil {
		slog.Error("user not found", slog.String("package", "userservice"))
		return nil, errors.New("user not found")
	}

	user := response.UserResponse{
		ID:        userExists.ID,
		Name:      userExists.Name,
		Email:     userExists.Email,
		CreatedAt: userExists.CreatedAt,
		UpdatedAt: userExists.UpdatedAt,
	}

	return &user, nil
}

func (s *service) FindManyUsers(ctx context.Context) (*response.ManyUsersResponse, error) {
	findManyUsers, err := s.repository.FindManyUsers(ctx)
	if err != nil {
		slog.Error("error to find meny users", "err", err, slog.String("package", "userservice"))
		return nil, err
	}

	users := response.ManyUsersResponse{}
	for _, user := range findManyUsers {
		userReponse := response.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}

		users.Users = append(users.Users, userReponse)
	}

	return &users, nil
}

func (s *service) UpdateUserPassword(ctx context.Context, u *dto.UpdateUserPasswordDto, id string) error {
	userExists, err := s.repository.FindUserByID(ctx, id)
	if err != nil {
		slog.Error("error to search user by id", "err", err, slog.String("package", "userservice"))
		return err
	}

	if userExists == nil {
		slog.Error("user not found", slog.String("package", "userservice"))
		return errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userExists.Password), []byte(u.OldPassword))
	if err != nil {
		slog.Error("invlid password", slog.String("package", "userservice"))
		return errors.New("invalid password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userExists.Password), []byte(u.Password))
	if err != nil {
		slog.Error("new passowrd is equal to old password", slog.String("package", "userservice"))
		return errors.New("new passowrd is equal to old password")
	}

	passwordEncrypted, err := bcrypt.GenerateFromPassword([]byte(u.Password), 12)
	if err != nil {
		slog.Error("error to encrypt password", "err", err, slog.String("package", "userservice"))
		return errors.New("error to encrypt password")
	}

	err = s.repository.UpdatePassword(ctx, string(passwordEncrypted), id)
	if err != nil {
		slog.Error("error to update password", "err", err, slog.String("package", "userservice"))
		return err
	}

	return nil
}
