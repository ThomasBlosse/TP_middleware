package helpers

import (
	"Consumer/internal/models"
	"bytes"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

func FindingDifferentCollection(collections []models.Collection) {
	for _, collection := range collections {
		query := "http://localhost:8080/collections/" + collection.Uid
		resp, err := http.Get(query)
		if err != nil {
			logrus.Fatalf("Error while fetching collection: %s", err.Error())
		}
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusNotFound {
			creatingCollectionIfNotInTimetable(collection, query)
			continue
		}
		var body []byte
		body, err = io.ReadAll(resp.Body)
		if err != nil {
			logrus.Fatalf("Error while reading collection: %s", err.Error())
		}
		var timetable models.Collection
		if err := json.Unmarshal(body, &timetable); err != nil {
			logrus.Fatalf("Error while unmarshalling resources: %s", err.Error())
		}

	}

}

func creatingCollectionIfNotInTimetable(collection models.Collection, query string) {
	logrus.Infof("Collection not found, creating new collection")
	body, err := json.Marshal(collection)
	if err != nil {
		logrus.Fatalf("Error while marshalling collection: %s", err.Error())
	}

	createResp, err := http.Post(query, "application/json", bytes.NewBuffer(body))
	if err != nil {
		logrus.Fatalf("Error while creating collection: %s", err.Error())
	}
	defer createResp.Body.Close()

	if createResp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(createResp.Body)
		logrus.Fatalf("Failed to create collection. Status: %d - Response: %s", createResp.StatusCode, string(body))
	}

}
