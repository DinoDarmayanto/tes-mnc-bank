package model

import (
	"time"
)

type Merchant struct {
	Id            int       `json:"id"`
	Name          string    `json:"name"`
	AccountNumber string    `json:"account_number"`
	RegisteredAt  time.Time `json:"registered_at"`
}
