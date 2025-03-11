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

	// parsing datas in object slice
	collections := []models.Collection{}
	for rows.Next() {
		var data models.Collection
		err = rows.Scan(&data.ResourceIds, &data.Uid, &data.Description, &data.Name, &data.Started, &data.End, &data.Location, &data.LastUpdate)
		if err != nil {
			return nil, err
		}
		collections = append(collections, data)
	}
	// don't forget to close rows
	_ = rows.Close()

	return collections, err
}

func GetCollectionById(id uuid.UUID) (*models.Collection, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	row := db.QueryRow("SELECT * FROM collections WHERE id=?", id.String())
	helpers.CloseDB(db)

	var collection models.Collection
	err = row.Scan(&collection.ResourceIds, &collection.Uid, &collection.Description, &collection.Name, &collection.Started, &collection.End, &collection.Location, &collection.LastUpdate)
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

	_, err = db.Exec("INSERT INTO collections (resourceIds, uid, description, name, started, end, location, lastupdate) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
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

func PutCollectionById(collectionId uuid.UUID, start time.Time, end time.Time, location string) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	defer helpers.CloseDB(db)

	_, err = db.Exec("UPDATE collections SET started=?, end=?, location=?, lastupdate=? WHERE collection_id=?",
		start,
		end,
		location,
		time.Now(),
		collectionId.String(),
	)

	if err != nil {
		return err
	}

	return nil
}

func DeleteCollectionById(collectionId uuid.UUID) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM collections WHERE collection_id=?", collectionId.String())

	if err != nil {
		return err
	}

	return nil
}
