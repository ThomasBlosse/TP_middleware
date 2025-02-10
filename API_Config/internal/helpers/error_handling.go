package helpers

import (
	"API_Config/internal/models"
	"API_Config/internal/services/resource_service"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
)

func CheckResourceExists(resourceId uuid.UUID) error {
	_, err := resource_service.GetResourceById(resourceId)
	if err != nil {
		if customErr, ok := err.(*models.CustomError); ok {
			return customErr
		}
		logrus.Errorf("error retrieving resource: %s", err.Error())
		return &models.CustomError{
			Message: "Resource not found",
			Code:    http.StatusNotFound,
		}
	}
	return nil
}
