package resources

import (
	"API_Config/internal/models"
	"API_Config/internal/services/resources/service"
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

func GetAlerts(w http.ResponseWriter, _ *http.Request) {
	resources, err := service.GetAllResources()
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
	body, _ := json.Marshal(resources)
	_, _ = w.Write(body)
}
