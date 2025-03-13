package main

import (
	"Scheduler/internal/helpers"
	"Scheduler/internal/models"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	resp, err := http.Get("http://localhost:8080/resources")
	if err != nil {
		logrus.Fatalf("Error while fetching resources: %s", err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		logrus.Fatalf("Unexpected status code: %d - Response: %s", resp.StatusCode, string(body))
	}
	var body []byte
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		logrus.Fatalf("Error while reading resources: %s", err.Error())
	}

	var resources []models.Resources
	if err := json.Unmarshal(body, &resources); err != nil {
		logrus.Fatalf("Error while unmarshalling resources: %s", err.Error())
	}

	var resourceToFetch strings.Builder

	for i, resource := range resources {
		if i > 0 {
			resourceToFetch.WriteString(",")
		}
		uidStr := strconv.Itoa(resource.Uid)
		resourceToFetch.WriteString(uidStr)
	}
	formattedResources := resourceToFetch.String()

	// Retrieve data from EDT
	url := fmt.Sprintf("https://edt.uca.fr/jsp/custom/modules/plannings/anonymous_cal.jsp?resources=%s&projectId=2&calType=ical&nbWeeks=8&displayConfigId=128", formattedResources)

	resp, err = http.Get(url)
	if err != nil {
		logrus.Fatalf("Error while fetching calendar data: %s", err.Error())
	}

	// Read all and store in value
	rawData, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Fatalf("Error while reading calendar data: %s", err.Error())
	}
	var eventArray []map[string]string
	eventArray, err = helpers.ParseICalEvents(rawData)
	if err != nil {
		logrus.Fatalf("Error parsing calendar: %s", err)
	}

	var collections []models.Collection
	collections, err = helpers.ConvertEventsToCollections(eventArray)
	if err != nil {
		logrus.Fatalf("Error converting events: %s", err)
	}

	jsonData, err := json.Marshal(collections)
	if err != nil {
		logrus.Fatalf("Error while marshalling collections: %s", err.Error())
	}

}
