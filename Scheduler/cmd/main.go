package main

import (
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
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
}
