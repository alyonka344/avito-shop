package model

import "github.com/gofrs/uuid/v5"

type MerchItem struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Price int       `json:"price"`
}
