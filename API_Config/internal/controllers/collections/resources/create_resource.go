package resources

import (
	"API_Config/internal/models"
	"API_Config/internal/services/collections/resource_service"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

func CreateResource(w http.ResponseWriter, r *http.Request) {
	var newResource models.Resources

	if err := json.NewDecoder(r.Body).Decode(&newResource); err != nil {
		logrus.Errorf("error decoding request body: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := resource_service.PostResource(newResource)
	if err != nil {
		logrus.Errorf("error adding resource: %s", err.Error())
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
	body, _ := json.Marshal(newResource)
	_, _ = w.Write(body)
}
