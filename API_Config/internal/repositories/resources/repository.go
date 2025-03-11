package resources

import (
	"API_Config/internal/helpers"
	"API_Config/internal/models"

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
		var tempId uuid.UUID
		err := rows.Scan(&resource.Name, &resource.Uid, &tempId)
		if err != nil {
			return nil, err
		}
		resources = append(resources, resource)
	}
	_ = rows.Close()

	return resources, nil
}

func GetResourceByUid(uid uuid.UUID) (*models.Resources, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	row := db.QueryRow("SELECT * FROM resources WHERE uid = ?", uid.String())
	helpers.CloseDB(db)

	var resource models.Resources
	var tempId uuid.UUID
	err = row.Scan(&resource.Name, &resource.Uid, &tempId)
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

	Id, _ := uuid.NewV4()

	_, err = db.Exec("INSERT INTO resources (name, uid, id) VALUES (?, ?, ?)",
		resource.Name,
		resource.Uid,
		Id.String(),
	)

	if err != nil {
		return err
	}

	helpers.CloseDB(db)

	return nil
}

func DeleteResourceById(resourceId uuid.UUID) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	_, err = db.Exec("DELETE FROM resources  WHERE id=?", resourceId.String())
	helpers.CloseDB(db)
	return err
}
