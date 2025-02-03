package collections

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"API_Timetable/internal/services/collections"
	"net/http"
)

// DeleteCollectionItem
// @Tags         collections
// @Summary      Delete an item from a collection.
// @Description  Remove a specific item from a collection.
// @Param        id       path      string  true  "Collection UUID formatted ID"
// @Param        item_id  path      string  true  "Item UUID formatted ID"
// @Success      200      "Item deleted successfully"
// @Failure      500      "Something went wrong"
// @Router       /collections/{id}/items/{item_id} [delete]
func DeleteCollectionItem(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	collectionId, _ := ctx.Value("collectionId").(uuid.UUID)
	itemId, _ := ctx.Value("itemId").(uuid.UUID)

	// Call the service to delete the item from the collection
	if err := collections.DeleteCollectionById(collectionId, itemId); err != nil {
		logrus.Errorf("error deleting item: %s", err.Error())
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
	body, _ := json.Marshal(map[string]string{"message": "Item deleted successfully"})
	_, _ = w.Write(body)
}
