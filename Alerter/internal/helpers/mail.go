package helpers

import (
	"Alerter/config"
	"Alerter/internal/models"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func SendMail(notification models.Notification) {
	alerts := getAlerts(notification.ResourceIds)
	for _, alert := range alerts {
		err := writeMail(alert.Email, notification.Description, notification.OldValue, notification.NewValue)
		if err != nil {
			logrus.Error(err)
		}
	}
}

func getAlerts(ResourceIds []int) []models.Alerts {
	var allAlerts []models.Alerts
	alertsMap := make(map[string]struct{})

	for _, resourceId := range ResourceIds {
		resp, err := http.Get("http://localhost:8080/alerts/" + strconv.Itoa(resourceId))
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

func writeMail(email string, description string, base string, change string) error {
	mailContent, mailSubject, err := config.GetStringFromEmbeddedTemplate("templates/email.html", models.MailTemplateData{
		Description: description,
		Base:        base,
		Change:      change,
	})
	if err != nil {
		logrus.Fatalf("Error parsing email template: %s", err.Error())
	}

	mailReq := models.MailRequest{
		Recipient: email,
		Subject:   mailSubject,
		Content:   mailContent,
	}
	reqBody, err := json.Marshal(mailReq)
	if err != nil {
		return fmt.Errorf("error marshalling mail request: %w", err)
	}

	req, err := http.NewRequest("POST", "https://mail-api.edu.forestier.re/mail", bytes.NewBuffer(reqBody))
	if err != nil {
		return fmt.Errorf("error creating HTTP request: %w", err)
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Erreur lors du chargement du fichier .env : %v", err)
	}
	apiToken := os.Getenv("API_TOKEN")

	req.Header.Set("Authorization", apiToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending mail: %w", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("failed to send email, status: %d, response: %s", resp.StatusCode, string(body))
	}

	logrus.Infof("Email successfully sent to %s", email)
	return nil
}
