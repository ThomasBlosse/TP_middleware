package resources

import (
	"API_Config/internal/models"
	service "API_Config/internal/services/resources"
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

// CreateResource
// @Summary Create resource
// @Description Creates a new resource
// @Tags resources
// @Accept json
// @Produce json
// @Param resource body models.Resources true "New resource"
// @Success 201 {object} models.Resources
// @Failure 400 {object} models.CustomError
// @Failure 500 {object} models.CustomError
// @Router /resources [post]
func CreateResource(w http.ResponseWriter, r *http.Request) {
	var newResource models.Resources

	if err := json.NewDecoder(r.Body).Decode(&newResource); err != nil {
		logrus.Errorf("error decoding request body: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := service.PostResource(newResource)
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
