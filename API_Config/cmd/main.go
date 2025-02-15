package main

import (
	"API_Config/internal/controllers/alerts"
	"API_Config/internal/controllers/resources"
	"net/http"

	"github.com/sirupsen/logrus"
)

func main() {
	r := chi.NewRouter()

	r.Route("/alerts", func(r chi.Router) {
		r.Get("/", alerts.GetAlerts)
		r.Post("/", alerts.CreateAlert)

		r.Route("/{id}", func(r chi.Router) {
			r.Use(alerts.Ctx)
			r.Get("/", alerts.GetAlert)
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
			r.Get("/", resources.GetResourceUid)
			r.Delete("/", resources.DeleteResource)
		})
	})

	logrus.Info("[INFO] Web server started. Now listening on *:8080")
	logrus.Fatalln(http.ListenAndServe(":8080", r))

}
