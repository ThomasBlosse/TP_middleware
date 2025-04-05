package resources

import (
	"API_Config/internal/models"
	service "API_Config/internal/services/resources"
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

// GetResource
// @Summary Get resource by ID
// @Description Retrieves a resource given its ID
// @Tags resources
// @Produce json
// @Success 200 {object} models.Resources
// @Failure 500 {object} models.CustomError
// @Router /resources/{id} [get]
func GetResource(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ucaId, _ := ctx.Value("resourceId").(int)
	resource, err := service.GetResourceByUid(ucaId)
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
	body, _ := json.Marshal(resource)
	_, _ = w.Write(body)
}
