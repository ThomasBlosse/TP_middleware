package alerts

import (
	"API_Config/internal/models"
	service "API_Config/internal/services/alerts"
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

// GetAlerts
// @Summary Get all alerts
// @Description Retrieves all the alerts in the system
// @Tags alerts
// @Produce json
// @Success 200 {array} models.Alerts
// @Failure 500 {object} models.CustomError
// @Router /alerts [get]
func GetAlerts(w http.ResponseWriter, _ *http.Request) {
	// calling service
	alerts, err := service.GetAllAlerts()
	if err != nil {
		logrus.Errorf("error : %s", err.Error())
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

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(alerts)
	_, _ = w.Write(body)
}
