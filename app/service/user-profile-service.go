package service

import (
	"os"

	"github.com/go-resty/resty/v2"
)

// GetPushNotificationID : sends a message to the device
func GetPushNotificationID(username string) (string, error) {
	client := resty.New()
	request := client.R()

	pushNotification := struct {
		ID string `json:"id"`
	}{}
	request.SetResult(&pushNotification)

	_, err := request.Get(os.Getenv("API_USER_PROFILE_BASE_URL") + "/" + username + "/push-notification")
	if err != nil {
		return "", err
	}

	return pushNotification.ID, nil
}
