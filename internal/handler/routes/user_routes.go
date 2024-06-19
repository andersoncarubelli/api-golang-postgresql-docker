package routes

import (
	"github.com/andersoncarubelli/api-golang-postgresql-docker/internal/handler/userhandler"
	"github.com/go-chi/chi"
)

func InitUserRoutes(router chi.Router, h userhandler.UserHandler) {
	router.Route("/user", func(r chi.Router) {
		r.Post("/", h.CreateUser)
		r.Patch("/{id}", h.UpdateUser)
		r.Get("/{id}", h.GetUserByID)
		r.Delete("/{id}", h.DeleteUser)
		r.Get("/", h.FindManyUsers)
		r.Patch("/password/{id}", h.UpdateUserPassword)
	})
	router.Route("/auth", func(r chi.Router) {
		r.Post("/login", h.Login)
	})
}
