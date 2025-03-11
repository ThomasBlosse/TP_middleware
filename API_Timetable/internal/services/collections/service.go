package collections

import (
	"API_Timetable/internal/models"
	repository "API_Timetable/internal/repositories/collections"
	"database/sql"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
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

func GetCollectionByUid(uid string) (*models.Collection, error) {
	collection, err := repository.GetCollectionByUid(uid)
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
	err := repository.PostCollection(collection)
	if err != nil {
		logrus.Errorf("error adding collection: %s", err.Error())
		return &models.CustomError{
			Message: "Something went wrong",
			Code:    http.StatusInternalServerError,
		}
	}

	return nil
}

func PutCollectionById(collectionId uuid.UUID, start time.Time, end time.Time, location string) error {
	err := repository.PutCollectionById(collectionId, start, end, location)
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

func DeleteCollectionById(collectionId uuid.UUID) error {
	err := repository.DeleteCollectionById(collectionId)
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
