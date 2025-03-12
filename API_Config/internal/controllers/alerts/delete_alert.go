package alerts

import (
	"API_Config/internal/models"
	service "API_Config/internal/services/alerts"
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

func DeleteAlert(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	email, _ := ctx.Value("Email").(string)

	err := service.DeleteAlertByEmail(email)
	if err != nil {
		logrus.Errorf("error deleting alert: %s", err.Error())

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
	body, _ := json.Marshal(map[string]string{"message": "Alert deleted successfully"})
	_, _ = w.Write(body)
}
