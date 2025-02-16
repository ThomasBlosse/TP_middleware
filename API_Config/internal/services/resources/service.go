package resources

import (
	"API_Config/internal/models\"
	repository "API_Config/internal/repositories/resources"
	"database/sql"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
)

func GetAllResources() ([]models.Resources, error) {
	allResources, err := repository.GetAllResources()
	if err != nil {
		logrus.Errorf("error retrieving all the resources : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return allResources, nil
}

func GetResourceByUid(uid uuid.UUID) (*models.Resources, error) {
	resource, err := repository.GetResourceByUid(uid)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil, &models.CustomError{
				Message: "resource not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error retrieving resource : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return resource, nil
}

func PostResource(resource models.Resources) error {
	err := repository.PostResource(resource)
	if err != nil {
		logrus.Errorf("error adding resource: %s", err.Error())
		return &models.CustomError{
			Message: "Something went wrong",
			Code:    http.StatusInternalServerError,
		}
	}
	return nil
}

func DeleteResourceById(resourceId uuid.UUID) error {
	err := repository.DeleteResourceById(resourceId)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return &models.CustomError{
				Message: "Resource not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error deleting resource: %s", err.Error())
		return &models.CustomError{
			Message: "Something went wrong",
			Code:    http.StatusInternalServerError,
		}
	}

	return nil
}
