package main

import (
	"Scheduler/internal/models"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

func ConvertEventsToCollections(eventArray []map[string]string) ([]models.Collection, error) {
	var resourceMapping = map[string][]int{
		"M1 GROUPE 1 LANGUE": {13295},
		"M1 GROUPE 2 LANGUE": {13345},
		"M1 GROUPE 3 LANGUE": {13397},
		"M1 GROUPE 1 OPTION": {7224},
		"M1 GROUPE 2 OPTION": {7225},
		"M1 GROUPE 3 OPTION": {62962},
		"M1 GROUPE OPTION":   {62090},
		"M1 -- Tutorat L2":   {56529},
		"MASTER 1 INFO":      {13295, 13345, 13397, 7224, 7225, 62962, 62090, 56529},
	}

	var collections []models.Collection

	for _, event := range eventArray {
		started, err := time.Parse("20060102T150405Z", event["DTSTART"])
		if err != nil {
			logrus.Warnf("Invalid DTSTART format: %s", event["DTSTART"])
			continue
		}

		end, err := time.Parse("20060102T150405Z", event["DTEND"])
		if err != nil {
			logrus.Warnf("Invalid DTEND format: %s", event["DTEND"])
			continue
		}

		lastUpdate, err := time.Parse("20060102T150405Z", event["LAST-MODIFIED"])
		if err != nil {
			logrus.Warnf("Invalid LAST-MODIFIED format: %s", event["LAST-MODIFIED"])
			continue
		}

		var resourceIds []int
		description := event["DESCRIPTION"]

		for desc, resIds := range resourceMapping {
			if strings.Contains(description, desc) {
				for _, resId := range resIds {
					resUUID, err := uuid.Parse(resId)
					if err == nil {
						resourceIds = append(resourceIds, &resUUID)
					} else {
						logrus.Warnf("Invalid Resource UUID: %s", resId)
					}
				}
			}
		}

		collection := models.Collection{
			ResourceIds: resourceIds,
			Uid:         event["UID"],
			Description: description,
			Name:        event["SUMMARY"],
			Started:     started,
			End:         end,
			Location:    event["LOCATION"],
			LastUpdate:  lastUpdate,
		}

		collections = append(collections, collection)
	}

	return collections, nil
}
