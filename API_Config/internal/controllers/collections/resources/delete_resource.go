package resources

import (
	"API_Config/internal/models"
	"API_Config/internal/services/collections/resource_service"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

func DeleteResource(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	collectionId, _ := ctx.Value("collectionId").(uuid.UUID)

	err := resource_service.DeleteResourceById(collectionId)
	if err != nil {
		logrus.Errorf("error deleting resource: %s", err.Error())

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
	body, _ := json.Marshal(map[string]string{"message": "Resource deleted successfully"})
	_, _ = w.Write(body)
}
