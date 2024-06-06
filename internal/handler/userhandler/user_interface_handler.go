package userhandler

import (
	"net/http"

	"github.com/andersoncarubelli/api-golang-postgresql-docker/internal/service/userservice"
)

func NewUserHandler(service userservice.UserService) UserHandler {
	return &handler{
		service,
	}
}

type handler struct {
	service userservice.UserService
}

type UserHandler interface {
	CreateUser(w http.ResponseWriter, r *http.Request) error
}
