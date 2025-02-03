package collections

import (
	"database/sql"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	repository "middleware/example/internal/repositories/collections"
	"net/http"
)

func GetAllCollections() ([]models.Collection, error) {
	var err error
	// calling repository
	collections, err := repository.GetAllCollections()
	// managing errors
	if err != nil {
		logrus.Errorf("error retrieving collections : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return collections, nil
}

func GetCollectionById(id uuid.UUID) (*models.Collection, error) {
	collection, err := repository.GetCollectionById(id)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil, &models.CustomError{
				Message: "collection not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error retrieving collections : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return collection, err
}


func PostCollection(collection models.Collection) error {
	err := repository.AddCollection(collection)
	if err != nil {
		logrus.Errorf("error adding collection: %s", err.Error())
		return &models.CustomError{
			Message: "Something went wrong",
			Code:    http.StatusInternalServerError,
		}
	}

	return nil
}

func PutCollectionById(collectionId uuid.UUID, item models.Item) error {
	err := repository.PutCollectionById(collectionId, item)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return &models.CustomError{
				Message: "Item not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error updating item: %s", err.Error())
		return &models.CustomError{
			Message: "Something went wrong",
			Code:    http.StatusInternalServerError,
		}
	}

	return nil
}


func DeleteCollectionById(collectionId uuid.UUID, itemId uuid.UUID) error {
	err := repository.DeleteCollectionById(collectionId, itemId)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return &models.CustomError{
				Message: "Item not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error deleting item: %s", err.Error())
		return &models.CustomError{
			Message: "Something went wrong",
			Code:    http.StatusInternalServerError,
		}
	}

	return nil
}
