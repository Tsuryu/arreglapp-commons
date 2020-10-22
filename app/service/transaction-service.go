package service

import (
	"os"

	"github.com/Tsuryu/arreglapp-commons/app/models"
	"github.com/go-resty/resty/v2"
)

// NewTransaction : send a new request to transactions
func NewTransaction(transaction *models.Transaction, withSecurityCode bool) error {
	client := resty.New()

	request := client.R()
	request.SetBody(transaction)
	if withSecurityCode {
		request.SetHeader("with-security-code", "true")
	}
	request.SetResult(&transaction)

	_, err := request.Post(os.Getenv("API_TRANSACTION_BASE_URL") + "/transaction")
	return err
}
