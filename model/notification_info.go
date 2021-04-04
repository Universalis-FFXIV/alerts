package model

// NotificationInfo describes a notification and its target user.
type NotificationInfo struct {
	TargetUser   string       `json:"targetUser"`
	Notification Notification `json:"notification"`
}
