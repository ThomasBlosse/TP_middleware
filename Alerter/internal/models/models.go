package models

type Notification struct {
	ResourceIds []int  `json:"resourceIds"`
	Description string `json:"description"`
	OldValue    string `json:"oldvalue"`
	NewValue    string `json:"newvalue"`
}

type Alerts struct {
	Email   string   `json:"email"`
	Targets []string `json:"targets"`
}
