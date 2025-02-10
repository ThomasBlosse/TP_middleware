package collections

import (
	"API_Timetable/internal/helpers"
	"API_Timetable/internal/models"
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
		err := rows.Scan(&alert.Name, &alert.Uid, &alert.Id)
		if err != nil {
			return nil, err
		}
		alerts = append(alerts, alert)
	}
	_ = rows.Close()

	return alerts, nil
}
