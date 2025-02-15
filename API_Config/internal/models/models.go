package models

import "github.com/gofrs/uuid"

type Resources struct {
	Name string     `json:"name"`
	Uid  string     `json:"UCA_ID"`
	Id   *uuid.UUID `json:"id"`
}

type Alerts struct {
	Email   string      `json:"email"`
	Targets interface{} `json:"targets"`
	Id      *uuid.UUID  `json:"id"`
}
