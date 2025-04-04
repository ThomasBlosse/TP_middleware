package helpers

import (
	"Alerter/internal/models"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strconv"
)

func SendMail(notification models.Notification) {
	alerts := getAlerts(notification.ResourceIds)
}

func getAlerts(ResourceIds []int) []models.Alerts {
	var allAlerts []models.Alerts
	alertsMap := make(map[string]struct{})

	for _, resourceId := range ResourceIds {
		resp, err := http.Get("http://localhost:8081/alerts/" + strconv.Itoa(resourceId))
		if err != nil {
			logrus.Fatalf("Error while fetching alerts: %s", err.Error())
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			body, _ := io.ReadAll(resp.Body)
			logrus.Fatalf("Unexpected status code: %d - Response: %s", resp.StatusCode, string(body))
		}
		var body []byte
		body, err = io.ReadAll(resp.Body)
		if err != nil {
			logrus.Fatalf("Error while reading alerts : %s", err.Error())
		}
		var alerts []models.Alerts
		if err := json.Unmarshal(body, &alerts); err != nil {
			logrus.Fatalf("Error while unmarshalling resources: %s", err.Error())
		}

		for _, alert := range alerts {
			if _, exists := alertsMap[alert.Email]; !exists {
				alertsMap[alert.Email] = struct{}{}
				allAlerts = append(allAlerts, alert)
			}
		}
	}
	return allAlerts
}
