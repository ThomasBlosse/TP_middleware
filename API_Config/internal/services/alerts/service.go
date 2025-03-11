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

	resourceIds, err := checkingTargets(targets)
	if err != nil {
		return err
	}

	err = checkingIfAllResourcesExist(resourceIds)
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

	resourceIds, err := checkingTargets(newTargets)
	if err != nil {
		return err
	}

	err = checkingIfAllResourcesExist(resourceIds)
	if err != nil {
		return err
	}

	err = repository.PutAlert(email, newTargets)
	if err != nil {
		logrus.Errorf("error updating alert: %s", err.Error())
		return &models.CustomError{
			Message: "Something went wrong",
			Code:    http.StatusInternalServerError,
		}
	}

	return nil
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
