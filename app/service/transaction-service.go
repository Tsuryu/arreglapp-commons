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

// GetTransactions : fetch transactions by username
func GetTransactions(username, reference string) ([]models.Transaction, error) {
	client := resty.New()
	request := client.R()
	transactions := []models.Transaction{}
	request.SetResult(&transactions)

	endpoint := "/transaction?"
	endpoint = addParam(endpoint, "username", username)
	endpoint = addParam(endpoint, "reference", reference)

	_, err := request.Get(os.Getenv("API_TRANSACTION_BASE_URL") + endpoint)
	return transactions, err
}

func addParam(endpoint, key, value string) string {
	if value != "" {
		if endpoint[len(endpoint)-1:] != "?" {
			endpoint += "&"
		}
		endpoint += key + "=" + value
	}
	return endpoint
}
