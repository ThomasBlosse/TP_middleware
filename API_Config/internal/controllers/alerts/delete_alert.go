package alerts

import (
	"API_Config/internal/models"
	"API_Config/internal/services/alerts/service"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

func DeleteAlert(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	collectionId, _ := ctx.Value("collectionId").(uuid.UUID)

	err := service.DeleteAlertById(collectionId)
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
