package alerts

import (
	"API_Config/internal/models"
	service "API_Config/internal/services/resources"
	"net/http"
	"strconv"

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

func checkingTargets(targets []string) ([]int, error) {
	if len(targets) != 1 {
		for _, target := range targets {
			if target == "all" {
				return nil, &models.CustomError{
					Message: "If 'all' is present, no other resources should be specified",
					Code:    http.StatusBadRequest,
				}
			}
		}
	}

	var resourceIds []int
	for _, target := range targets {
		resourceId, errConv := strconv.Atoi(target)
		if errConv != nil {
			logrus.Errorf("error converting target to int: %s", target)
			return nil, &models.CustomError{
				Message: "Something went wrong",
				Code:    http.StatusInternalServerError,
			}
		}
		resourceIds = append(resourceIds, resourceId)
	}

	return resourceIds, nil
}
