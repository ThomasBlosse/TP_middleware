package collections

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"API_Timetable/internal/models"
	"API_Timetable/internal/services/collections"
	"net/http"
)

// CreateCollection
// @Tags         collections
// @Summary      Create a new collection.
// @Description  Add a new collection to the database.
// @Param        collection   body      models.Collection  true  "Collection data"
// @Success      201          {object}  models.Collection
// @Failure      400          "Invalid request body"
// @Failure      500          "Something went wrong"
// @Router       /collections [post]
func CreateCollection(w http.ResponseWriter, r *http.Request) {
	var newCollection models.Collection

	// Decode the request body into the collection model
	if err := json.NewDecoder(r.Body).Decode(&newCollection); err != nil {
		logrus.Errorf("error decoding request body: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Generate a new UUID for the collection
	newCollection.Id = uuid.Must(uuid.NewV4())

	// Call the service to add the collection
	if err := collections.AddCollection(newCollection); err != nil {
		logrus.Errorf("error adding collection: %s", err.Error())
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

	// Respond with the created collection
	w.WriteHeader(http.StatusCreated)
	body, _ := json.Marshal(newCollection)
	_, _ = w.Write(body)
}
