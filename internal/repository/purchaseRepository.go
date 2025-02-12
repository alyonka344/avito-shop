package repository

import (
	"avito-shop/internal/model"
	"github.com/gofrs/uuid/v5"
)

type PurchaseRepository interface {
	Create(purchase *model.Purchase) error
	GetById(purchaseID uuid.UUID) (model.Purchase, error)
}
