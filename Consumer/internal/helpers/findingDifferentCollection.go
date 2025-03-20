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

func GeneratingNotification(collections []models.Collection) []models.Notification {
	var notifications []models.Notification
	for _, collection := range collections {
		different := false
		query := "http://localhost:8081/collections/" + collection.Uid
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
			different = true
		}

		notification = compareStart(collection, timetable)
		if notification.Description != "" {
			notifications = append(notifications, notification)
			different = true
		}

		if different {
			updateCollection(collection, query)
		}
	}
	return notifications
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
			Description: "Changement de Salle : " + collection.Name,
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
			Description: "Changement de début du cours" + collection.Name,
			OldValue:    formatDateTime(timetable.Started),
			NewValue:    formatDateTime(collection.Started),
		}
	}
	return models.Notification{}
}

func formatDateTime(t time.Time) string {
	return fmt.Sprintf("Le %02d/%02d à %02d:%02d", t.Day(), t.Month(), t.Hour(), t.Minute())
}

func updateCollection(collection models.Collection, query string) {
	body, err := json.Marshal(collection)
	if err != nil {
		logrus.Errorf("error marshalling collection: %w", err)
	}

	req, err := http.NewRequest("PUT", query, bytes.NewBuffer(body))
	if err != nil {
		logrus.Errorf("error creating PUT request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logrus.Errorf("error sending PUT request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		logrus.Errorf("PUT request failed with status %d: %s", resp.StatusCode, body)
	}
}
