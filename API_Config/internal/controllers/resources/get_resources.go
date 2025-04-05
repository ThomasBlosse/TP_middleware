package resources

import (
	"API_Config/internal/models"
	service "API_Config/internal/services/resources"
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

// GetResources
// @Summary Get all resources
// @Description Retrieves all the resources
// @Tags resources
// @Produce json
// @Success 200 {array} models.Resources
// @Failure 500 {object} models.CustomError
// @Router /resources [get]
func GetResources(w http.ResponseWriter, _ *http.Request) {
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
