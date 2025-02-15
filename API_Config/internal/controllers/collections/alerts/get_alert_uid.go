package alerts

import (
	"API_Config/internal/models"
	"API_Config/internal/services/collections/alert_service"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

func GetAlertUid(w http.ResponseWriter, r *http.Request) {
	ucaIdParam := chi.URLParam(r, "uid")
	ucaId, err := uuid.FromString(ucaIdParam)
	if err != nil {
		logrus.Errorf("parsing error : %s", err.Error())
		customError := &models.CustomError{
			Message: fmt.Sprintf("cannot parse uid (%s) as UUID", ucaIdParam),
			Code:    http.StatusUnprocessableEntity,
		}
		w.WriteHeader(customError.Code)
		body, _ := json.Marshal(customError)
		_, _ = w.Write(body)
		return
	}

	alert, err := alert_service.GetAlertsByResource(ucaId)
	if err != nil {
		logrus.Errorf("error : %s", err.Error())
		customError, isCustom := err.(*models.CustomError)
		if isCustom {
			w.WriteHeader(customError.Code)
			body, _ := json.Marshal(customError)
			_, _ = w.Write(body)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			body, _ := json.Marshal(models.CustomError{
				Message: "An internal server error occurred.",
				Code:    http.StatusInternalServerError,
			})
			_, _ = w.Write(body)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(alert)
	_, _ = w.Write(body)
}
