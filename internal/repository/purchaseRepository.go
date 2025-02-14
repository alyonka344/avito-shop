package repository

import (
	"avito-shop/internal/model"
)

type PurchaseRepository interface {
	Create(purchase *model.Purchase) error
	GetAllByUserName(userName string) ([]model.Purchase, error)
}
