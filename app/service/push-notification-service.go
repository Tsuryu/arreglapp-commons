package service

import (
	"os"

	"github.com/Tsuryu/arreglapp-commons/app/models"
	"github.com/go-resty/resty/v2"
)

// SendPushNotification : sends a message to the device
func SendPushNotification(notification models.PushNotification) error {
	client := resty.New()
	request := client.R()

	notification.Notification.ClickAction = "FLUTTER_NOTIFICATION_CLICK"
	request.SetBody(notification)

	_, err := request.Post(os.Getenv("API_FIREBASE_URL"))
	return err
}
