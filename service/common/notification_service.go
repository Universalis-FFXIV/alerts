package common

import "github.com/Universalis-FFXIV/alerts/model"

type NotificationService interface {
	SendNotification(targetUser string, notification *model.Notification) error
}
