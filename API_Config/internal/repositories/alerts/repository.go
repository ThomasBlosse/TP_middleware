package alerts

import (
	"API_Config/internal/helpers"
	"API_Config/internal/models"
	"strconv"
	"strings"

	"github.com/gofrs/uuid"
)

func GetAllAlerts() ([]models.Alerts, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT email, targets FROM alerts")
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	alerts := []models.Alerts{}
	for rows.Next() {
		var alert models.Alerts
		var targetsJSON string
		err := rows.Scan(&alert.Email, &targetsJSON)
		if err != nil {
			return nil, err
		}

		alert.Targets = strings.Split(targetsJSON, ",")
		alerts = append(alerts, alert)
	}
	_ = rows.Close()

	return alerts, nil
}

func GetAlertsByResource(resourceId int) ([]models.Alerts, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	defer helpers.CloseDB(db)
	var alerts []models.Alerts

	query := "SELECT email, targets FROM alerts WHERE targets LIKE ?"
	rows, err := db.Query(query, "%"+strconv.Itoa(resourceId)+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var alert models.Alerts
		var targetsJSON string
		err := rows.Scan(&alert.Email, &targetsJSON)
		if err != nil {
			return nil, err
		}

		alerts = append(alerts, alert)
	}

	rows, err = db.Query("SELECT email, targets FROM alerts WHERE targets = 'all'")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var alert models.Alerts
		var targetsJSON string
		err := rows.Scan(&alert.Email, &targetsJSON)
		if err != nil {
			return nil, err
		}

		alerts = append(alerts, alert)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return alerts, nil
}

func PostAlert(alert models.Alerts) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}

	targetsJSON := strings.Join(alert.Targets, ",")
	Id, _ := uuid.NewV4()

	_, err = db.Exec("INSERT INTO alerts (email, targets, id) VALUES (?, ?, ?)",
		alert.Email,
		targetsJSON,
		Id,
	)

	if err != nil {
		return err
	}

	helpers.CloseDB(db)

	return nil
}

func PutAlert(email string, newTargets []string) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}

	targetsJSON := strings.Join(newTargets, ",")

	_, err = db.Exec("UPDATE alerts SET targets = ? WHERE email = ?",
		targetsJSON,
		email,
	)
	helpers.CloseDB(db)
	return err
}

func DeleteAlertByEmail(email string) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	_, err = db.Exec("DELETE FROM resources  WHERE email = ?", email)
	helpers.CloseDB(db)
	return err
}
