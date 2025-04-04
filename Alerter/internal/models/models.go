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

type MailTemplateData struct {
	EventName   string
	Description string
	Base        string
	Change      string
}

type MailRequest struct {
	Recipient string `json:"recipient"`
	Subject   string `json:"subject"`
	Content   string `json:"content"`
}
