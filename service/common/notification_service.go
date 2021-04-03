package common

type NotificationService interface {
	SendNotification(targetUser string, body string) error
}
