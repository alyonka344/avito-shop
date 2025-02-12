package repository

import (
	"avito-shop/internal/model"
	"github.com/gofrs/uuid/v5"
)

type TransactionRepository interface {
	Create(transaction *model.Transaction) error
	GetAllByUserId(userID uuid.UUID) ([]model.Transaction, error)
}
