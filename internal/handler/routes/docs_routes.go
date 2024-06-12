package routes

import (
	"github.com/andersoncarubelli/api-golang-postgresql-docker/docs/custom"
	"github.com/go-chi/chi"
	_ "github.com/swaggo/http-swagger/example/go-chi/docs"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

var (
	docsURL = "http://localhost:8080/docs/doc.json"
)

// @title Swagger Dark Mode
// @version 1.0
// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization
func InitDocsRoutes(r chi.Router) {
	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL(docsURL),
		httpSwagger.AfterScript(custom.CustomLayoutJS),
		httpSwagger.DocExpansion("none"),
		httpSwagger.UIConfig(map[string]string{
			"defaultModelsExpandDepth": `"-1"`,
		}),
	))
}
