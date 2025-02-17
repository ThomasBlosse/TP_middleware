package main

import (
	"API_Config/internal/controllers/alerts"
	"API_Config/internal/controllers/resources"
	"API_Config/internal/helpers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

func main() {
	r := chi.NewRouter()

	r.Route("/alerts", func(r chi.Router) {
		r.Get("/", alerts.GetAlerts)
		r.Post("/", alerts.CreateAlert)

		r.Route("/{id}", func(r chi.Router) {
			r.Use(alerts.Ctx)
			r.Get("/", alerts.GetAlertsUid)
			r.Put("/", alerts.UpdateAlert)
			r.Delete("/", alerts.DeleteAlert)
		})
	})

	r.Route("/resources", func(r chi.Router) {
		r.Get("/", resources.GetResources)
		r.Post("/", resources.CreateResource)

		r.Route("/{id}", func(r chi.Router) {
			r.Use(resources.Ctx)
			r.Get("/", resources.GetResource)
			r.Delete("/", resources.DeleteResource)
		})
	})

	logrus.Info("[INFO] Web server started. Now listening on *:8080")
	logrus.Fatalln(http.ListenAndServe(":8080", r))

}

func init() {
	db, err := helpers.OpenDB()
	if err != nil {
		logrus.Fatalf("error while opening database : %s", err.Error())
	}
	schemes := []string{
		`CREATE TABLE IF NOT EXISTS ressources (
    		name TEXT NOT NULL,
    		uid VARCHAR(255) NOT NULL,
			id UUID PRIMARY KEY NOT NULL UNIQUE
    	);`,
	}
	for _, scheme := range schemes {
		if _, err := db.Exec(scheme); err != nil {
			logrus.Fatalln("Could not generate table resources ! Error was : " + err.Error())
		}
	}

	schemes = []string{
		`CREATE TABLE IF NOT EXISTS alerts (
    		email VARCHAR(255) NOT NULL,
    		targets TEXT NOT NULL,
			id UUID PRIMARY KEY NOT NULL UNIQUE
    	);`,
	}
	for _, scheme := range schemes {
		if _, err := db.Exec(scheme); err != nil {
			logrus.Fatalln("Could not generate table alerts ! Error was : " + err.Error())
		}
	}

	helpers.CloseDB(db)
}
