package model

import (
	"time"
)

type Transaction struct {
	ID              int       `json:"id"`
	CustomerID      int       `json:"customer_id"`
	MerchantID      int       `json:"merchant_id"`
	Amount          float64   `json:"amount"`
	TransactionTime time.Time `json:"transaction_time"`
}
