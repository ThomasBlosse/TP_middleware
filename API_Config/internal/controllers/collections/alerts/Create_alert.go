package alerts

import (
	"API_Config/internal/models"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

func CreateAlert(w http.ResponseWriter, r *http.Request) {
	var newAlert models.Alerts

	if err := json.NewDecoder(r.Body).Decode(&newAlert); err != nil {
		logrus.Errorf("error decoding request body: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := alert_service.PostAlert(newAlert)
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
	body, _ := json.Marshal(newCollection)
	_, _ = w.Write(body)
}
