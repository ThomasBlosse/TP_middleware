package alerts

import (
	"API_Config/internal/models"
	service "API_Config/internal/services/alerts"
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

// CreateAlert
// @Summary Create a new alert
// @Description Creates a new alert from the JSON payload and saves it in the database
// @Tags alerts
// @Accept  json
// @Produce  json
// @Param alert body models.Alerts true "Alert to create"
// @Success 201 {object} models.Alerts
// @Failure 400 {object} models.CustomError "Invalid JSON"
// @Failure 500 {object} models.CustomError "Internal Server Error"
// @Router /alerts [post]
func CreateAlert(w http.ResponseWriter, r *http.Request) {
	var newAlert models.Alerts

	if err := json.NewDecoder(r.Body).Decode(&newAlert); err != nil {
		logrus.Errorf("error decoding request body: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := service.PostAlert(newAlert)
	if err != nil {
		logrus.Errorf("error adding alert: %s", err.Error())
		customError, isCustom := err.(*models.CustomError)
		if isCustom {
			w.WriteHeader(customError.Code)
			body, _ := json.Marshal(customError)
			_, _ = w.Write(body)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	body, _ := json.Marshal(newAlert)
	_, _ = w.Write(body)
}
