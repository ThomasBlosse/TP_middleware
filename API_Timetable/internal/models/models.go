package models

import (
	"github.com/gofrs/uuid"
	"time"
)

type Collection struct {
	Id      		*uuid.UUID 		`json:"id"`
	ResourceIds		[]*uuid.UUID		`json:"resourceIds"`
	Uid 			string 			`json:"uid"`
	Description 		string 			`json:"description"`
	Name 			string 			`json:"name"`
	Started			time.Time		`json:"started"`
	End 			time.Time		`json:"end"`
	Location		str			`json:"Location"`
	LastUpdate		time.Time		`json:"lastupdate"`
}
