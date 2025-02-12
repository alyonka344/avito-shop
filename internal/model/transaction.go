package model

import (
	"github.com/gofrs/uuid/v5"
	"time"
)

type Transaction struct {
	ID                uuid.UUID `json:"id"`
	FromUserID        int       `json:"from_user_id"`
	ToUserID          int       `json:"to_user_id"`
	Amount            int       `json:"amount"`
	TransactionStatus Status    `json:"transaction_status"`
	CreatedAt         time.Time `json:"created_at"`
}
