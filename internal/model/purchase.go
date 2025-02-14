package model

import (
	"github.com/gofrs/uuid/v5"
	"time"
)

type Purchase struct {
	ID        uuid.UUID `json:"id"  db:"id"`
	UserName  string    `json:"username"  db:"username"`
	MerchName string    `json:"merch_name"  db:"merch_name"`
	CreatedAt time.Time `json:"created_at"  db:"created_at"`
}
