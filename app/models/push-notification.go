package models

// PushNotification : push notification structure
type PushNotification struct {
	To           string           `json:"to,omitempty"`
	Notification NotificationInfo `json:"notification,omitempty"`
	Data         interface{}      `json:"data,omitempty"`
}

// NotificationInfo : push notification info
type NotificationInfo struct {
	Title       string `json:"title,omitempty"`
	Body        string `json:"body,omitempty"`
	ClickAction string `json:"click_action,omitempty"`
}
