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
func GetTransactions(username string) ([]models.Transaction, error) {
	client := resty.New()
	request := client.R()
	transactions := []models.Transaction{}
	request.SetResult(&transactions)

	_, err := request.Get(os.Getenv("API_TRANSACTION_BASE_URL") + "/transactions?username=" + username)
	return transactions, err
}
