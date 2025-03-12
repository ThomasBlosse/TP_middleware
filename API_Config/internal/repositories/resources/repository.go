package resources

import (
	"API_Config/internal/helpers"
	"API_Config/internal/models"
	"strconv"

	"github.com/gofrs/uuid"
)

func GetAllResources() ([]models.Resources, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT name, uid FROM resources")
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	resources := []models.Resources{}
	for rows.Next() {
		var resource models.Resources
		err := rows.Scan(&resource.Name, &resource.Uid)
		if err != nil {
			return nil, err
		}
		resources = append(resources, resource)
	}
	_ = rows.Close()

	return resources, nil
}

func GetResourceByUid(uid int) (*models.Resources, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	row := db.QueryRow("SELECT name, uid FROM resources WHERE uid = ?", strconv.Itoa(uid))
	helpers.CloseDB(db)

	var resource models.Resources
	err = row.Scan(&resource.Name, &resource.Uid)
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

func DeleteResourceByUid(uid int) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	_, err = db.Exec("DELETE FROM resources  WHERE uid=?", strconv.Itoa(uid))
	helpers.CloseDB(db)
	return err
}
