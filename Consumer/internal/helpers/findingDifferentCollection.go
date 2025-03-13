package helpers

import (
	"Consumer/internal/models"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"time"
)

func FindingDifferentCollection(collections []models.Collection) {
	var notifications []models.Notification
	for _, collection := range collections {
		query := "http://localhost:8080/collections/" + collection.Uid
		resp, err := http.Get(query)
		if err != nil {
			logrus.Fatalf("Error while fetching collection: %s", err.Error())
		}
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusNotFound {
			creatingCollection(collection)
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
		notification := compareLocation(collection, timetable)
		if notification.Description != "" {
			notifications = append(notifications, notification)
		}

		notification = compareStart(collection, timetable)
		if notification.Description != "" {
			notifications = append(notifications, notification)
		}

	}

}

func creatingCollection(collection models.Collection) {
	logrus.Infof("Collection not found, creating new collection")
	body, err := json.Marshal(collection)
	if err != nil {
		logrus.Errorf("Error while marshalling collection: %s", err.Error())
	}

	createResp, err := http.Post("http://localhost:8080/collections", "application/json", bytes.NewBuffer(body))
	if err != nil {
		logrus.Errorf("Error while creating collection: %s", err.Error())
	}
	defer createResp.Body.Close()

	if createResp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(createResp.Body)
		logrus.Errorf("Failed to create collection. Status: %d - Response: %s", createResp.StatusCode, string(body))
	}

}

func compareLocation(collection models.Collection, timetable models.Collection) models.Notification {
	if collection.Location != timetable.Location {
		return models.Notification{
			ResourceIds: collection.ResourceIds,
			Description: "Changement de Salle",
			OldValue:    timetable.Location,
			NewValue:    collection.Location,
		}
	}
	return models.Notification{}
}

func compareStart(collection models.Collection, timetable models.Collection) models.Notification {
	if collection.Started != timetable.Started {
		return models.Notification{
			ResourceIds: collection.ResourceIds,
			Description: "Changement de début du cours",
			OldValue:    formatDateTime(timetable.Started),
			NewValue:    formatDateTime(collection.Started),
		}
	}
	return models.Notification{}
}

func formatDateTime(t time.Time) string {
	return fmt.Sprintf("Le %02d/%02d à %02d:%02d", t.Day(), t.Month(), t.Hour(), t.Minute())
}
