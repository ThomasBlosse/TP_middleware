package models

import "github.com/gofrs/uuid"

type Resources struct {
	Name string `json:"name"`
	Uid  int    `json:"uid"`
}

type Alerts struct {
	Email   string      `json:"email"`
	Targets interface{} `json:"targets"`
	Id      *uuid.UUID  `json:"id"`
}
