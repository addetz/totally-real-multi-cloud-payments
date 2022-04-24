package models

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/google/uuid"

)

type PaymentResponse struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}

type PaymentRequest struct {
	TransactionID string `json:"tx_id"`
	UserID        string `json:"user_id"`
	ClientID      string `json:"client_id"`
	FromAcct      string `json:"from_acct"`
	ToAcct        string `json:"to_acct"`
	Amount        int    `json:"amount"`
}

func NewPaymentRequest(req *http.Request) (*PaymentRequest, error) {
	var payment PaymentRequest
	if err := json.NewDecoder(req.Body).Decode(&payment); err != nil {
		return nil, fmt.Errorf("error during payment decode:%s", err.Error())
	}
	payment.TransactionID = uuid.New().String()
	return &payment, nil
}
