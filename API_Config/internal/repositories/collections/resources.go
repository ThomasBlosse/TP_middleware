package collections

import (
	"API_Timetable/internal/helpers"
	"API_Timetable/internal/models"
)

func GetAllResources() ([]models.Resources, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM resources")
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	resources := []models.Resources{}
	for rows.Next() {
		var resource models.Resources
		err := rows.Scan(&resource.Name, &resource.Uid, &resource.Id)
		if err != nil {
			return nil, err
		}
		resources = append(resources, resource)
	}
	_ = rows.Close()

	return resources, nil
}
