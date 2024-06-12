package main

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/andersoncarubelli/api-golang-postgresql-docker/config/env"
	"github.com/andersoncarubelli/api-golang-postgresql-docker/config/logger"
	"github.com/andersoncarubelli/api-golang-postgresql-docker/internal/database"
	"github.com/andersoncarubelli/api-golang-postgresql-docker/internal/database/sqlc"
	"github.com/andersoncarubelli/api-golang-postgresql-docker/internal/handler/routes"
	"github.com/andersoncarubelli/api-golang-postgresql-docker/internal/handler/userhandler"
	"github.com/andersoncarubelli/api-golang-postgresql-docker/internal/repository/userrepository"
	"github.com/andersoncarubelli/api-golang-postgresql-docker/internal/service/userservice"
	"github.com/go-chi/chi"
)

func main() {
	logger.InitLogger()
	slog.Info("starting api")

	_, err := env.LoadingConfig(".")
	if err != nil {
		slog.Error("failed to load environment variables", err, slog.String("package", "main"))
		return
	}

	dbConnection, err := database.NewDBConnection()
	if err != nil {
		slog.Error("error to connect to database", "err", err, slog.String("package", "main"))
		return
	}

	router := chi.NewRouter()
	queries := sqlc.New(dbConnection)

	// user
	newUserRepository := userrepository.NewUserRepository(dbConnection, queries)
	newUserService := userservice.NewUserService(newUserRepository)
	newUserHandler := userhandler.NewUserHandler(newUserService)

	// init routes
	routes.InitUserRoutes(router, newUserHandler)
	routes.InitDocsRoutes(router)

	port := fmt.Sprintf(":%s", env.Env.GoPort)
	slog.Info(fmt.Sprintf("server running on port %s", port))
	err = http.ListenAndServe(port, router)
	if err != nil {
		slog.Error("error to start server", err, slog.String("package", "main"))
	}

}
