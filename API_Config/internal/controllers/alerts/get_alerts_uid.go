package alerts

import (
	"API_Config/internal/models"
	service "API_Config/internal/services/alerts"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

func GetAlertsUid(w http.ResponseWriter, r *http.Request) {
	ucaIdParam := chi.URLParam(r, "targets")
	ucaId, err := uuid.FromString(ucaIdParam)
	if err != nil {
		logrus.Errorf("parsing error : %s", err.Error())
		customError := &models.CustomError{
			Message: fmt.Sprintf("cannot parse targets (%s) as UUID", ucaIdParam),
			Code:    http.StatusUnprocessableEntity,
		}
		w.WriteHeader(customError.Code)
		body, _ := json.Marshal(customError)
		_, _ = w.Write(body)
		return
	}

	alerts, err := service.GetAlertsByResource(ucaId)
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
	body, _ := json.Marshal(alerts)
	_, _ = w.Write(body)
}
