package alerts

import (
	"API_Config/internal/models"
	repository "API_Config/internal/repositories/alerts"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

func GetAllAlerts() ([]models.Alerts, error) {
	allAlerts, err := repository.GetAllAlerts()
	if err != nil {
		logrus.Errorf("error retrieving all the alerts : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return allAlerts, nil
}

func GetAlertsByResource(resourceId int) ([]models.Alerts, error) {
	if err := checkResourceExists(resourceId); err != nil {
		return nil, err
	}

	resourceAlerts, err := repository.GetAlertsByResource(resourceId)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil, &models.CustomError{
				Message: "repository not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error retrieving alert: %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    http.StatusInternalServerError,
		}
	}
	return resourceAlerts, nil
}

func PostAlert(alert models.Alerts) error {
	targets := alert.Targets

	if len(targets) != 1 {
		for _, target := range targets {
			if target == "all" {
				return &models.CustomError{
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
			return &models.CustomError{
				Message: "Something went wrong",
				Code:    http.StatusInternalServerError,
			}
		}
		resourceIds = append(resourceIds, resourceId)
	}

	err := checkingIfAllResourcesExist(resourceIds)
	if err != nil {
		return err
	}

	err = repository.PostAlert(alert)
	if err != nil {
		logrus.Errorf("error adding alert: %s", err.Error())
		return &models.CustomError{
			Message: "Something went wrong",
			Code:    http.StatusInternalServerError,
		}
	}
	return nil
}

func PutAlert(email string, newTargets []string) error {

	if len(newTargets) != 1 {
		for _, target := range newTargets {
			if target == "all" {
				return &models.CustomError{
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
			return &models.CustomError{
				Message: "Something went wrong",
				Code:    http.StatusInternalServerError,
			}
		}
		resourceIds = append(resourceIds, resourceId)
	}

	err := checkingIfAllResourcesExist(resourceIds)
	if err != nil {
		return err
	}

	return UpdateAlert(email, newTargets)
}

func DeleteAlertById(alertId uuid.UUID) error {
	err := repository.DeleteAlertByEmail(alertId)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return &models.CustomError{
				Message: "Alert not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error deleting alert: %s", err.Error())
		return &models.CustomError{
			Message: "Something went wrong",
			Code:    http.StatusInternalServerError,
		}
	}

	return nil
}
