package collections

import (
	"API_Timetable/internal/models"
	"API_Timetable/internal/services/collections"
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

// GetCollections
// @Tags         collections
// @Summary      Get collections.
// @Description  Get collections.
// @Success      200            {array}  models.Collection
// @Failure      500             "Something went wrong"
// @Router       /collections [get]
func GetCollections(w http.ResponseWriter, _ *http.Request) {
	// calling service
	collections, err := collections.GetAllCollections()
	if err != nil {
		// logging error
		logrus.Errorf("error : %s", err.Error())
		customError, isCustom := err.(*models.CustomError)
		if isCustom {
			// writing http code in header
			w.WriteHeader(customError.Code)
			// writing error message in body
			body, _ := json.Marshal(customError)
			_, _ = w.Write(body)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(collections)
	_, _ = w.Write(body)
}
