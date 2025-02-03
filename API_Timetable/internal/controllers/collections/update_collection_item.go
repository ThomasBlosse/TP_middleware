package collections

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"API_Timetable/internal/models"
	"API_Timetable/internal/services/collections"
	"net/http"
)

// UpdateCollectionItem
// @Tags         collections
// @Summary      Update an item in a collection.
// @Description  Modify an existing item in a collection.
// @Param        id          path      string             true  "Collection UUID formatted ID"
// @Param        item        body      models.Item        true  "Updated item data"
// @Success      200         "Item updated successfully"
// @Failure      400         "Invalid request body"
// @Failure      500         "Something went wrong"
// @Router       /collections/{id}/items [put]
func UpdateCollectionItem(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	collectionId, _ := ctx.Value("collectionId").(uuid.UUID)

	var updatedItem models.Item
	// Decode the request body into the updated item model
	if err := json.NewDecoder(r.Body).Decode(&updatedItem); err != nil {
		logrus.Errorf("error decoding request body: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Call the service to update the item in the collection
	if err := collections.UpdateItem(collectionId, updatedItem); err != nil {
		logrus.Errorf("error updating item: %s", err.Error())
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

	// Respond with success message
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"message": "Item updated successfully"}`))
}
