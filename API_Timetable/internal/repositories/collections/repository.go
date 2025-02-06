package collections

import (
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"

	"github.com/gofrs/uuid"
)

func GetAllCollections() ([]models.Collection, error) 
{
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
		err = rows.Scan(&data.Id, &data.Content)
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
	err = row.Scan(&collection.Id, &collection.Content)
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

	_, err = db.Exec("INSERT INTO collections (id, resourceIds, uid, description, name, started, end, location, lastupdate) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
		collection.Id.String(),
		collection.ResourceIds.String(),
		collection.Uid,
		collection.Description,
		collection.Name,
		collection.Started,
		collection.End,
		collection.Location,
		collection.LastUpdate,
	)

	if err != nil {
		return nil, err
	}

	return &collection, nil
}



func PutCollectionById(collectionId uuid.UUID, item models.Item) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	defer helpers.CloseDB(db)

	_, err = db.Exec("UPDATE items SET name=?, description=? WHERE collection_id=? AND id=?",
		item.Name,
		item.Description,
		collectionId.String(),
		item.Id.String(),
	)

	if err != nil {
		return err
	}

	return nil
}



func DeleteCollectionById(collectionId uuid.UUID, itemId uuid.UUID) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM items WHERE collection_id=? AND id=?", collectionId.String(), itemId.String())

	if err != nil {
		return err
	}

	return nil
}
