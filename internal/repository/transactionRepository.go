package repository

import (
	"avito-shop/internal/model"
	"github.com/gofrs/uuid/v5"
)

type TransactionRepository interface {
	Create(transaction *model.Transaction) error
	GetAllSentByUserId(userID uuid.UUID) ([]model.Transaction, error)
	GetAllReceivedByUserId(userID uuid.UUID) ([]model.Transaction, error)
}
