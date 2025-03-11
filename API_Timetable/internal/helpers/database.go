package helpers

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/gofrs/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
)

func OpenDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "file:collections.db")

	if err != nil {
		fmt.Printf("%s", err.Error())
		db.SetMaxOpenConns(1)
	}
	return db, err
}
func CloseDB(db *sql.DB) {
	err := db.Close()
	if err != nil {
		logrus.Errorf("error closing db : %s", err.Error())
	}
}

func UUIDSliceToString(uuids []*uuid.UUID) string {

	var strUUIDs []string

	for _, u := range uuids {

		strUUIDs = append(strUUIDs, u.String())

	}

	return strings.Join(strUUIDs, ",")

}

func StringToUUIDSlice(s string) ([]*uuid.UUID, error) {
	parts := strings.Split(s, ",")
	var uuids []*uuid.UUID

	for _, part := range parts {
		part = strings.TrimSpace(part)
		uuid, err := uuid.FromString(part)
		if err != nil {
			return nil, err
		}
		uuids = append(uuids, &uuid)
	}
	return uuids, nil
}
