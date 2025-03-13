package collections

import (
	"API_Timetable/internal/helpers"
	"API_Timetable/internal/models"
	"strconv"
	"strings"
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
		var collection models.Collection
		var tempId uuid.UUID
		var resourceIdsJson string
		err = rows.Scan(&tempId, &resourceIdsJson, &collection.Uid, &collection.Description, &collection.Name, &collection.Started, &collection.End, &collection.Location, &collection.LastUpdate)
		if err != nil {
			return nil, err
		}
		resourceIdsStr := strings.Split(resourceIdsJson, ",")
		for _, resourceId := range resourceIdsStr {
			resource, err := strconv.Atoi(resourceId)
			if err != nil {
				return nil, err
			}
			collection.ResourceIds = append(collection.ResourceIds, resource)
		}
		collections = append(collections, collection)
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
	var resourceIdsJson string
	err = row.Scan(&tempId, &resourceIdsJson, &collection.Uid, &collection.Description, &collection.Name, &collection.Started, &collection.End, &collection.Location, &collection.LastUpdate)
	if err != nil {
		return nil, err
	}
	resourceIds := strings.Split(resourceIdsJson, ",")
	collection.ResourceIds = make([]int, len(resourceIds))
	return &collection, err
}

func PostCollection(collection models.Collection) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	defer helpers.CloseDB(db)

	var strResourceIds []string
	for _, id := range collection.ResourceIds {
		strResourceIds = append(strResourceIds, strconv.Itoa(id))
	}

	resourceIdsJson := strings.Join(strResourceIds, ",")
	Id, _ := uuid.NewV4()

	_, err = db.Exec("INSERT INTO collections (id, resourceIds, uid, description, name, started, end, location, lastupdate) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
		Id,
		resourceIdsJson,
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
