package helpers

import (
	"database/sql"
	"fmt"
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
