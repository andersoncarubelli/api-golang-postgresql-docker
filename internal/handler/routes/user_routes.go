package routes

import (
	"github.com/andersoncarubelli/api-golang-postgresql-docker/config/env"
	"github.com/andersoncarubelli/api-golang-postgresql-docker/internal/handler/userhandler"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
)

func InitUserRoutes(router chi.Router, h userhandler.UserHandler) {
	router.Post("/user", h.CreateUser)
	router.Route("/user", func(r chi.Router) {
		r.Use(jwtauth.Verifier(env.Env.TokenAuth))
		r.Use(jwtauth.Authenticator)

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
