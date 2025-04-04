package helpers

import (
	"Scheduler/internal/models"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

func ConvertEventsToCollections(eventArray []map[string]string) ([]models.Collection, error) {
	var resourceMapping = map[string][]int{
		"M1 GROUPE 1 langue": {13295},
		"M1 GROUPE 2 langue": {13345},
		"M1 GROUPE 3 langue": {13397},
		"M1 Groupe 1 option": {7224},
		"M1 Groupe 2 option": {7225},
		"M1 Groupe 3 option": {62962},
		"M1 Groupe option":   {62090},
		"M1 -- Tutorat L2":   {56529},
		"MASTER 1 INFO":      {13295, 13345, 13397, 7224, 7225, 62962, 62090, 56529},
		"L3 S6":              {26888},
		"L3 info":            {26888},
		"LICENCE 3 Info":     {26888},
		"gA":                 {57454},
		"gB":                 {57455},
		"gC":                 {57456},
		"gD":                 {57457},
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
				resourceIds = append(resourceIds, resIds...)
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
