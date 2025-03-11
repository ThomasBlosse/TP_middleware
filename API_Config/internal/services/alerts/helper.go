package alerts

import (
	"API_Config/internal/models"
	service "API_Config/internal/services/resources"
	"net/http"

	"github.com/sirupsen/logrus"
)

func checkResourceExists(resourceId int) error {
	_, err := service.GetResourceByUid(resourceId)
	if err != nil {
		if customErr, ok := err.(*models.CustomError); ok {
			return customErr
		}
		logrus.Errorf("error retrieving resource: %s", err.Error())
		return &models.CustomError{
			Message: "Resource not found",
			Code:    http.StatusBadRequest,
		}
	}
	return nil
}

func checkingIfAllResourcesExist(targetsIds []int) error {
	for _, target := range targetsIds {
		err := checkResourceExists(target)
		if err != nil {
			return err
		}
	}
	return nil
}
