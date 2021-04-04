package model

type Notification struct {
	ItemName string `json:"itemName"`
	PageURL  string `json:"pageUrl"`
	Body     string `json:"body"`
}
