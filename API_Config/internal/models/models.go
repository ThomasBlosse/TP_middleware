package models

type Resources struct {
	Name string     `json:"name"`
	Uid  string     `json:"UCA ID"`
	Id   *uuid.UUID `json:"id"`
}
