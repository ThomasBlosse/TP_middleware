package alerts

import (
	"API_Config/internal/models"
	service "API_Config/internal/services/alerts"
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

// UpdateAlert
// @Summary Update alert
// @Description Updates the targets of an existing alert for the authenticated user
// @Tags alerts
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string "Alert updated successfully"
// @Failure 400 {object} models.CustomError
// @Failure 500 {object} models.CustomError
// @Router /alerts/{id} [put]
func UpdateAlert(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	email, _ := ctx.Value("Email").(string)

	var updatedAlert models.Alerts

	if err := json.NewDecoder(r.Body).Decode(&updatedAlert); err != nil {
		logrus.Errorf("error decoding request body: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := service.PutAlert(email, updatedAlert.Targets)
	if err != nil {
		logrus.Errorf("error updating alert: %s", err.Error())
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

	// Respond with success message
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"message": "Alert updated successfully"}`))
}
