package resources

import (
	"API_Config/internal/models"
	service "API_Config/internal/services/resources"
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

// DeleteResource
// @Summary Delete resource
// @Description Deletes a resource by its ID
// @Tags resources
// @Success 200 {object} map[string]string "Resource deleted successfully"
// @Failure 500 {object} models.CustomError
// @Router /resources/{id} [delete]
func DeleteResource(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ucaId, _ := ctx.Value("resourceId").(int)

	err := service.DeleteResourceByUid(ucaId)
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
