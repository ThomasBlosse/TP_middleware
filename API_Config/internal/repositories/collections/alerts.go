package collections

import (
	"API_Timetable/internal/helpers"
	"API_Timetable/internal/models"
	"encoding/json"
)

func GetAllAlerts() ([]models.Alerts, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM alerts")
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	alerts := []models.Alerts{}
	for rows.Next() {
		var alert models.Alerts
		err := rows.Scan(&alert.Email, &alert.Targets, &alert.Id)
		if err != nil {
			return nil, err
		}
		alerts = append(alerts, alert)
	}
	_ = rows.Close()

	return alerts, nil
}

func GetAlertById(id uuid.UUID) (*models.Alerts, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	row := db.QueryRow("SELECT * FROM alerts WHERE id=?", id.String())
	helpers.CloseDB(db)

	var alert models.Alerts
	err := rows.Scan(&alert.Email, &alert.Targets, &alert.Id)
	if err != nil {
		return nil, err
	}
	return &alert, nil
}

func GetAlertsByResource(resourceId uuid.UUID) ([]Alerts, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT email, targets, id FROM alerts")
	if err != nil {
		return nil, err
	}
	helpers.CloseDB(db)

	var alerts []Alerts
	for rows.Next() {
		var alert Alerts
		var targetsJSON string

		err := rows.Scan(&alert.Email, &targetsJSON, &alert.Id)
		if err != nil {
			return nil, err
		}

		var targetsMap map[string]interface{}
		if err := json.Unmarshal([]byte(targetsJSON), &targetsMap); err != nil {
			return nil, err
		}

		if allValue, exists := targetsMap["all"]; exists && allValue == true {
			alerts = append(alerts, alert)
		} else if resources, exists := targetsMap["resources"]; exists {
			if resourceList, ok := resources.([]interface{}); ok {
				for _, resource := range resourceList {
					if resStr, ok := resource.(string); ok {
						if resStr == resourceId.String() {
							alerts = append(alerts, alert)
							break
						}
					}
				}
			}
		}
	}
	rows.Close()

	return alerts, nil
}
