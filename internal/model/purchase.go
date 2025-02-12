package model

import (
	"github.com/gofrs/uuid/v5"
	"time"
)

type Purchase struct {
	ID        uuid.UUID `json:"id"`
	UserID    int       `json:"user_id"`
	MerchName string    `json:"merch_name"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
}
