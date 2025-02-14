package repository

import (
	"avito-shop/internal/model"
)

type TransactionRepository interface {
	Create(transaction *model.Transaction) error
	GetAllSentByUserName(userName string) ([]model.Transaction, error)
	GetAllReceivedByUserName(userName string) ([]model.Transaction, error)
}
