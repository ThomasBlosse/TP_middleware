package main

import (
	"encoding/json"
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

}
