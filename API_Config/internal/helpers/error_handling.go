package helpers

import (
	"API_Config/internal/models"
	service "API_Config/internal/services/resources"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
)

func CheckResourceExists(resourceId uuid.UUID) error {
	_, err := service.GetResourceById(resourceId)
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

func CheckingIfAllResourcesExist(targetsMap map[string]interface{}) error {
	if resources, exists := targetsMap["resources"]; exists {
		if resourceList, ok := resources.([]interface{}); ok {
			for _, resourceID := range resourceList {
				resourceIDStr, ok := resourceID.(string)
				if !ok {
					return &models.CustomError{
						Message: "Invalid resource ID format",
						Code:    http.StatusBadRequest,
					}
				}

				resourceUUID, err := uuid.FromString(resourceIDStr)
				if err != nil {
					return &models.CustomError{
						Message: "Invalid UUID format",
						Code:    http.StatusBadRequest,
					}
				}

				if err := CheckResourceExists(resourceUUID); err != nil {
					return err
				}
			}
		} else {
			return &models.CustomError{
				Message: "Invalid resources format",
				Code:    http.StatusBadRequest,
			}
		}
	}
	return nil
}
