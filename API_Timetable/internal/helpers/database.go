package helpers

import (
	"database/sql"
	"strings"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

func OpenDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "file:collections.db")
	if err != nil {
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
