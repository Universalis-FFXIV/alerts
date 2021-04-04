package common

import "github.com/Universalis-FFXIV/alerts/model"

// NotificationService handles communicating any notification to users.
type NotificationService interface {
	// SendNotification sends the provided notification to the provided user.
	SendNotification(targetUser string, notification *model.Notification) error
}
