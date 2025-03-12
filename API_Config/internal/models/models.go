package models

type Resources struct {
	Name string `json:"name"`
	Uid  int    `json:"uid"`
}

type Alerts struct {
	Email   string   `json:"email"`
	Targets []string `json:"targets"`
}
