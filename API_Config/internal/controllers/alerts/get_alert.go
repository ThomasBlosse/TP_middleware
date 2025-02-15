package alerts

import (
	"API_Config/internal/models"
	"API_Config/internal/services/alerts/service"
	"encoding/json"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

func GetAlert(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	collectionId, _ := ctx.Value("collectionId").(uuid.UUID)

	collection, err := service.GetAlertById(collectionId)
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
	body, _ := json.Marshal(collection)
	_, _ = w.Write(body)
}
