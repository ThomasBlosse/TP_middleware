package helpers

import (
	"Consumer/internal/models"
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
		if resp.StatusCode != http.StatusOK {
			body, _ := io.ReadAll(resp.Body)
			logrus.Fatalf("Unexpected status code: %d - Response: %s", resp.StatusCode, string(body))
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
