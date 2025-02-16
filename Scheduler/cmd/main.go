package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Fatalf("Error while reading resources: %s", err.Error())
	}

	var resources []models.Resource
	if err := json.Unmarshal(body, &resources); err != nil {
		logrus.Fatalf("Error while unmarshalling resources: %s", err.Error())
	}

	var resourceToFetch strings.Builder

	for i, resource := range resources {
		if i > 0 {
			resourceToFetch.WriteString(", ")
		}
		resourceToFetch.WriteString(resource.Uid)
	}
	formattedResources := resourceToFetch.String()

	// Retrieve data from EDT
	url := fmt.Sprintf("https://edt.uca.fr/jsp/custom/modules/plannings/anonymous_cal.jsp?resources=%s&projectId=2&calType=ical&nbWeeks=8&displayConfigId=128", formattedResources)

	resp, err = http.Get(url)
	if err != nil {
		logrus.Fatalf("Error while fetching calendar: %s", err.Error())
	}

	// Read all and store in value
	rawData, err := io.ReadAll(resp.Body)
	// TODO manage error

	// Create a line-reader from data
	scanner := bufio.NewScanner(bytes.NewReader(rawData))

	// Create vars
	var eventArray []map[string]string
	currentEvent := map[string]string{}

	currentKey := ""
	currentValue := ""

	inEvent := false

	// Inspect each line
	for scanner.Scan() {
		// Ignore calendar lines
		if !inEvent && scanner.Text() != "BEGIN:VEVENT" {
			continue
		}
		// If new event, go to next line
		if scanner.Text() == "BEGIN:VEVENT" {
			inEvent = true
			continue
		}

		// TODO if end event

		// TODO if multi-line data

		// Split scan
		fmt.Println(scanner.Text())
		splitted := strings.SplitN(scanner.Text(), ":", 2)
		currentKey = splitted[0]
		currentValue = splitted[1]

		// Store current event attribute
		currentEvent[currentKey] = currentValue
	}

	// TODO Transform to proper custom object

	// TODO parse to JSON and display

}
