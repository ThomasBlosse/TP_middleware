package main

import (
	"Scheduler/internal/helpers"
	"Scheduler/internal/models"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

func main() {

	helpers.Schedule()
	formattedResources := helpers.GetRessource()

	url := fmt.Sprintf("https://edt.uca.fr/jsp/custom/modules/plannings/anonymous_cal.jsp?resources=%s&projectId=2&calType=ical&nbWeeks=8&displayConfigId=128", formattedResources)

	resp, err := http.Get(url)
	if err != nil {
		logrus.Fatalf("Error while fetching calendar data: %s", err.Error())
	}

	rawData, err := io.ReadAll(resp.Body)
	if err != nil {

		logrus.Fatalf("Error while reading calendar data: %s", err.Error())
	}

	eventArray, err := helpers.ParseICalEvents(rawData)
	if err != nil {
		logrus.Fatalf("Error parsing calendar: %s", err)
	}

	collections, err := helpers.ConvertEventsToCollections(eventArray)
	if err != nil {
		logrus.Fatalf("Error converting events: %s", err)
	}

	jsonData, err := json.Marshal(collections)
	if err != nil {
		logrus.Fatalf("Error while marshalling collections: %s", err.Error())
	}
	var testData []models.Collection
	err = json.Unmarshal(jsonData, &testData)

	err = helpers.SendCollection(jsonData)
	if err != nil {
		logrus.Fatalf("Error while sending collections: %s", err.Error())
	}

}
