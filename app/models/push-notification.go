package models

// PushNotification : push notification structure
type PushNotification struct {
	To           string
	Notification NotificationInfo
	Data         interface{}
}

// NotificationInfo : push notification info
type NotificationInfo struct {
	Title       string
	Body        string
	ClickAction string
}
