package internal

import (
	_ "github.com/ThomasBlosse/TP_middleware/API_Config/docs" // swag init ira chercher les fichiers ici
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

// @title Config API
// @version 1.0
// @description APi to handle configurations
// @termsOfService http://swagger.io/terms/

// @contact.name Thomas Blosse
// @contact.email Thomas.BLOSSE@etu.uca.fr

// @host localhost:8080
// @BasePath /

func RegisterSwaggerRoutes(r chi.Router) {
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))
}
