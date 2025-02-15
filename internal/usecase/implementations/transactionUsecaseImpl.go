package implementations

import (
	"avito-shop/internal/model"
	"avito-shop/internal/repository"
	"errors"
	"github.com/gofrs/uuid/v5"
	"time"
)

type TransactionUsecase struct {
	userRepository        repository.UserRepository
	transactionRepository repository.TransactionRepository
}

func NewTransactionUseCase(
	userRepo repository.UserRepository,
	transactionRepo repository.TransactionRepository) *TransactionUsecase {
	return &TransactionUsecase{
		userRepository:        userRepo,
		transactionRepository: transactionRepo,
	}
}

func (t TransactionUsecase) TransferMoney(senderName string, recipientName string, amount int) error {
	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}

	transaction := model.Transaction{
		ID:        uuid.Must(uuid.NewV4()),
		FromUser:  senderName,
		ToUser:    recipientName,
		Amount:    amount,
		CreatedAt: time.Now(),
	}

	err := t.userRepository.Transfer(senderName, recipientName, amount)
	if err != nil {
		transaction.TransactionStatus = model.Failure
		err = t.transactionRepository.Create(&transaction)
		if err != nil {
			return err
		}
		return err
	}

	transaction.TransactionStatus = model.Success
	err = t.transactionRepository.Create(&transaction)
	if err != nil {
		return err
	}

	return nil
}
