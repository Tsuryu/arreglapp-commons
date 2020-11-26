package service

import (
	"errors"
	"os"

	"github.com/Tsuryu/arreglapp-commons/app/models"
	"github.com/go-resty/resty/v2"
)

// SendPushNotification : sends a message to the device
func SendPushNotification(notification models.PushNotification) error {
	apiKey := os.Getenv("API_KEY_FIREBASE")
	if apiKey == "" {
		return errors.New("Missing firebase apikey")
	}

	url := os.Getenv("API_FIREBASE_URL")
	if url == "" {
		return errors.New("Missing firebase url")
	}

	client := resty.New()
	request := client.R()

	notification.Notification.ClickAction = "FLUTTER_NOTIFICATION_CLICK"
	request.SetBody(notification)
	request.SetHeader("Authorization", apiKey)

	_, err := request.Post(url)
	return err
}
