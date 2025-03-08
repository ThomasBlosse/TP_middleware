package main

import (
	"API_Timetable/internal/controllers/collections"
	"API_Timetable/internal/helpers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

func main() {
	r := chi.NewRouter()

	r.Route("/collections", func(r chi.Router) {
		r.Get("/", collections.GetCollections)
		r.Post("/", collections.CreateCollection)

		r.Route("/{id}", func(r chi.Router) {
			r.Use(collections.Ctx)
			r.Get("/", collections.GetCollection)
			r.Put("/", collections.UpdateCollectionItem)
			r.Delete("/", collections.DeleteCollectionItem)
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
		`CREATE TABLE IF NOT EXISTS collections (
			id UUID PRIMARY KEY NOT NULL UNIQUE,
			resourceIds TEXT NOT NULL,
			uid VARCHAR(255) NOT NULL,
			description VARCHAR(255) NOT NULL,
			name VARCHAR(255) NOT NULL,
			started TIMESTAMP NOT NULL,
			end TIMESTAMP NOT NULL,
			location VARCHAR(255) NOT NULL,
			lastupdate TIMESTAMP NOT NULL
		);`,
	}
	for _, scheme := range schemes {
		if _, err := db.Exec(scheme); err != nil {
			logrus.Fatalln("Could not generate table ! Error was : " + err.Error())
		}
	}
	helpers.CloseDB(db)
}
