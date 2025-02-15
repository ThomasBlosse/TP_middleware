package alerts

import (
	"API_Config/internal/models"
	"API_Config/internal/services/collections/alert_service"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

func UpdateCollectionItem(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	collectionId, _ := ctx.Value("collectionId").(uuid.UUID)

	var updatedAlert models.Collection

	if err := json.NewDecoder(r.Body).Decode(&updatedAlert); err != nil {
		logrus.Errorf("error decoding request body: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := alert_service.PutAlert(collectionId, updatedAlert.Targets)
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
