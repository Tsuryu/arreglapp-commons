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

// GetTransaction : fetch transaction by id
func GetTransaction(traceID string) (models.Transaction, error) {
	client := resty.New()
	request := client.R()
	transaction := models.Transaction{}
	request.SetResult(&transaction)

	endpoint := "/transaction/" + traceID

	_, err := request.Get(os.Getenv("API_TRANSACTION_BASE_URL") + endpoint)
	return transaction, err
}

// GetTransactions : fetch transactions by username
func GetTransactions(transaction models.Transaction) ([]models.Transaction, error) {
	client := resty.New()
	request := client.R()
	transactions := []models.Transaction{}
	request.SetResult(&transactions)

	endpoint := "/transaction?"
	endpoint = addParam(endpoint, "username", transaction.Username)
	endpoint = addParam(endpoint, "reference", transaction.Reference)

	_, err := request.Get(os.Getenv("API_TRANSACTION_BASE_URL") + endpoint)
	return transactions, err
}

// AddTransactionDetail : add details to transaction
func AddTransactionDetail(detail models.TransactionDetail, traceID, securityCode string) (bool, error) {
	client := resty.New()

	request := client.R()
	request.SetBody(detail)
	request.SetHeader("security-code", securityCode)

	response, err := request.Post(os.Getenv("API_TRANSACTION_BASE_URL") + "/transaction/" + traceID + "/detail")
	return response.StatusCode() == 201, err
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
