package collections

import (
	"API_Config/internal/models\"
	alerts "API_Config/internal/repositories/collections"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
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
	_, err := resource_service.GetResourceById(resourceId)
	if err != nil {
		if customErr, ok := err.(*models.CustomError); ok {
			return nil, customErr
		}
		logrus.Errorf("error retrieving resource : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
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
