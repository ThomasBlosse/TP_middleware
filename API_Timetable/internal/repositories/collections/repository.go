package collections

import (
	"API_Timetable/internal/helpers"
	"API_Timetable/internal/models"
	"time"

	"github.com/gofrs/uuid"
)

func GetAllCollections() ([]models.Collection, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM collections")
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	collections := []models.Collection{}
	for rows.Next() {
		var data models.Collection
		var tempId uuid.UUID
		err = rows.Scan(&tempId, &data.ResourceIds, &data.Uid, &data.Description, &data.Name, &data.Started, &data.End, &data.Location, &data.LastUpdate)
		if err != nil {
			return nil, err
		}
		collections = append(collections, data)
	}
	_ = rows.Close()

	return collections, err
}

func GetCollectionByUid(uid string) (*models.Collection, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}

	row := db.QueryRow("SELECT * FROM collections WHERE uid=?", uid)
	helpers.CloseDB(db)

	var collection models.Collection
	var tempId uuid.UUID
	err = row.Scan(&tempId, &collection.ResourceIds, &collection.Uid, &collection.Description, &collection.Name, &collection.Started, &collection.End, &collection.Location, &collection.LastUpdate)
	if err != nil {
		return nil, err
	}
	return &collection, err
}

func PostCollection(collection models.Collection) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	defer helpers.CloseDB(db)

	Id, _ := uuid.NewV4()

	_, err = db.Exec("INSERT INTO collections (id, resourceIds, uid, description, name, started, end, location, lastupdate) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
		Id,
		collection.ResourceIds,
		collection.Uid,
		collection.Description,
		collection.Name,
		collection.Started,
		collection.End,
		collection.Location,
		collection.LastUpdate,
	)

	if err != nil {
		return err
	}

	return nil
}

func PutCollectionByUid(uid string, start time.Time, end time.Time, location string) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	defer helpers.CloseDB(db)

	_, err = db.Exec("UPDATE collections SET started=?, end=?, location=?, lastupdate=? WHERE uid=?",
		start,
		end,
		location,
		time.Now(),
		uid,
	)

	if err != nil {
		return err
	}

	return nil
}

func DeleteCollectionByUid(uid string) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM collections WHERE uid=?", uid)

	if err != nil {
		return err
	}

	return nil
}
