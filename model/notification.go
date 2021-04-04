package model

// Notification is the payload from which the user-facing notification will be constructed.
type Notification struct {
	ItemName string   `json:"itemName"`
	PageURL  string   `json:"pageUrl"`
	Reasons  []string `json:"reasons"`
}
