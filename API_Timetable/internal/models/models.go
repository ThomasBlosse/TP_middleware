package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Collection struct {
	Id          *uuid.UUID   `json:"id"`
	ResourceIds []*uuid.UUID `json:"resourceIds"`
	Uid         string       `json:"uid"`
	Description string       `json:"description"`
	Name        string       `json:"name"`
	Started     time.Time    `json:"started"`
	End         time.Time    `json:"end"`
	Location    str          `json:"location"`
	LastUpdate  time.Time    `json:"lastupdate"`
}
