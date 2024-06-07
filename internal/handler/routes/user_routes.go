package routes

import (
	"github.com/andersoncarubelli/api-golang-postgresql-docker/internal/handler/userhandler"
	"github.com/go-chi/chi"
)

func InitUserRoutes(router chi.Router, h userhandler.UserHandler) {
	router.Route("/user", func(r chi.Router) {
		r.Post("/", h.CreateUser)
	})
}
