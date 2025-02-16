package main

import (
	"net/http"
)

func main() {
	resp, err := http.Get("//localhost:8080/resources")
	if err != nil {
		logrus.Fatalf("Error while fetching resources: %s", err.Error())
	}
	defer resp.Body.Close()

}
