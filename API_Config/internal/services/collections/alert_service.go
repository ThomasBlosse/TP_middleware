package collections

import (
	"API_Config/internal/helpers"
	"API_Config/internal/models\"
	alerts "API_Config/internal/repositories/collections"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
)

func GetAllAlerts() ([]models.Alerts, error) {
	allAlerts, err := alerts.GetAllAlerts()
	if err != nil {
		logrus.Errorf("error retrieving all the alerts : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return allAlerts, nil
}

func GetAlertById(id uuid.UUID) (*models.Alerts, error) {
	alert, err := alerts.GetAlertById(id)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil, &models.CustomError{
				Message: "alert not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error retrieving alerts : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return alert, nil
}

func GetAlertsByResource(resourceId uuid.UUID) ([]models.Alerts, error) {
	if err := helpers.CheckResourceExists(resourceId); err != nil {
		return nil, err
	}

	resourceAlerts, err = alerts.GetAlertsByResource(resourceId)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil, &models.CustomError{
				Message: "alerts not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error retrieving alerts: %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    http.StatusInternalServerError,
		}
	}
	return resourceAlerts, nil
}

func InsertAlert(alert models.Alerts) error {
	err := alerts.PostAlert(alert)
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
			return insertAlert(alert)
		}
	}

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

				if err := helpers.CheckResourceExists(resourceUUID); err != nil {
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

	return InsertAlert(alert)
}

func UpdateAlert(alertId uuid.UUID, newTargets interface{}) error {
	err := alerts.PutAlert(alertId, newTargets)
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
			return UpdateAlert(alert)
		}
	}
}
