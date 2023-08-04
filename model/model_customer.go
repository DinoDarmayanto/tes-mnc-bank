package model

import (
	"time"
)

type Customer struct {
	ID           int       `json:"id"`
	Username     string    `json:"username"`
	Password     string    `json:"password"`
	Email        string    `json:"email"`
	RegisteredAt time.Time `json:"registered_at"`
}
