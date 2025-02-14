package model

import (
	"github.com/gofrs/uuid/v5"
	"time"
)

type Transaction struct {
	ID                uuid.UUID `json:"id" db:"id"`
	FromUser          string    `json:"from_user" db:"from_user"`
	ToUser            string    `json:"to_user" db:"to_user"`
	Amount            int       `json:"amount" db:"amount"`
	TransactionStatus Status    `json:"transaction_status" db:"transaction_status"`
	CreatedAt         time.Time `json:"created_at" db:"created_at"`
}
