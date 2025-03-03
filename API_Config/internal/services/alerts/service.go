package alerts

import (
	"API_Config/internal/models"
	repository "API_Config/internal/repositories/alerts"
	"database/sql"
	"net/http"

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

func GetAlertsByResource(resourceId uuid.UUID) ([]models.Alerts, error) {
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

func InsertAlert(alert models.Alerts) error {
	err := repository.PostAlert(alert)
	if err != nil {
		logrus.Errorf("error adding alert: %s", err.Error())
		return &models.CustomError{
			Message: "Something went wrong",
			Code:    http.StatusInternalServerError,
		}
	}
	return nil
}

func PostAlert(alert models.Alerts) error {
	targetsMap, ok := alert.Targets.(map[string]interface{})
	if !ok {
		return &models.CustomError{
			Message: "Invalid targets format",
			Code:    http.StatusBadRequest,
		}
	}

	if allValue, exists := targetsMap["all"]; exists {
		if all, ok := allValue.(bool); ok && all {
			if len(targetsMap) > 1 {
				return &models.CustomError{
					Message: "If 'all' is present, no other resources should be specified",
					Code:    http.StatusBadRequest,
				}
			}
			return InsertAlert(alert)
		}
	}

	err := checkingIfAllResourcesExist(targetsMap)
	if err != nil {
		return err
	}

	return InsertAlert(alert)
}

func UpdateAlert(alertId uuid.UUID, newTargets interface{}) error {
	err := repository.PutAlert(alertId, newTargets)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return &models.CustomError{
				Message: "Alert not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error updating alert: %s", err.Error())
		return &models.CustomError{
			Message: "Something went wrong",
			Code:    http.StatusInternalServerError,
		}
	}

	return nil
}

func PutAlert(alertId uuid.UUID, newTargets interface{}) error {
	targetsMap, ok := newTargets.(map[string]interface{})
	if !ok {
		return &models.CustomError{
			Message: "Invalid targets format",
			Code:    http.StatusBadRequest,
		}
	}

	if allValue, exists := targetsMap["all"]; exists {
		if all, ok := allValue.(bool); ok && all {
			if len(targetsMap) > 1 {
				return &models.CustomError{
					Message: "If 'all' is present, no other resources should be specified",
					Code:    http.StatusBadRequest,
				}
			}
			return UpdateAlert(alertId, newTargets)
		}
	}

	err := checkingIfAllResourcesExist(targetsMap)
	if err != nil {
		return err
	}

	return UpdateAlert(alertId, newTargets)
}

func DeleteAlertById(alertId uuid.UUID) error {
	err := repository.DeleteAlertById(alertId)
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
