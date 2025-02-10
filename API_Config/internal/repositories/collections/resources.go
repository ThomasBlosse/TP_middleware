package collections

import (
	"API_Timetable/internal/helpers"
	"API_Timetable/internal/models"
	"github.com/gofrs/uuid"
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

func GetResourceById(id uuid.UUID) (*models.Resources, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	row := db.QueryRow("SELECT * FROM resources WHERE id=?", id.String())
	helpers.CloseDB(db)

	var resource models.Resources
	err := rows.Scan(&resource.Name, &resource.Uid, &resource.Id)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

func PostResource(resource models.Resources) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO resources (name, uid, id) VALUES (?, ?, ?)",
		resource.Name,
		resource.Uid,
		resource.Id.String(),
	)

	if err != nil {
		return err
	}

	helpers.CloseDB(db)

	return nil
}
