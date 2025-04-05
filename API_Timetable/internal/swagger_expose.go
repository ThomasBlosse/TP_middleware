package internal

import (
	_ "github.com/ThomasBlosse/TP_middleware/API_Timetable/docs" // swag init ira chercher les fichiers ici
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

// @title Timetable API
// @version 1.0
// @description API to manage collections.
// @termsOfService http://swagger.io/terms/

// @contact.name Alban.Munari
// @contact.email Alban.MUNARI@etu.uca.fr

// @host localhost:8080
// @BasePath /

func RegisterSwaggerRoutes(r chi.Router) {
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))
}
