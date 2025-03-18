package models

import (
	"time"
)

type Collection struct {
	ResourceIds []int     `json:"resourceIds"`
	Uid         string    `json:"uid"`
	Description string    `json:"description"`
	Name        string    `json:"name"`
	Started     time.Time `json:"started"`
	End         time.Time `json:"end"`
	Location    string    `json:"location"`
	LastUpdate  time.Time `json:"lastupdate"`
}

type Resources struct {
	Name string `json:"name"`
	Uid  int    `json:"uid"`
}
