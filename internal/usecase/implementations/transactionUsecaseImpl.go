package implementations

import (
	"avito-shop/internal/model"
	"avito-shop/internal/repository"
	"avito-shop/internal/usecase"
	"errors"
	"github.com/gofrs/uuid/v5"
	"time"
)

type transactionUsecase struct {
	userRepository        repository.UserRepository
	transactionRepository repository.TransactionRepository
}

func NewTransactionUseCase(userRepo repository.UserRepository, transactionRepo repository.TransactionRepository) usecase.TransactionUsecase {
	return &transactionUsecase{
		userRepository:        userRepo,
		transactionRepository: transactionRepo,
	}
}

func (t transactionUsecase) TransferMoney(senderID uuid.UUID, recipientID uuid.UUID, amount int) error {
	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}

	transaction := model.Transaction{
		ID:         uuid.Must(uuid.NewV4()),
		FromUserID: senderID,
		ToUserID:   recipientID,
		Amount:     amount,
		CreatedAt:  time.Now(),
	}

	err := t.userRepository.Transfer(senderID, recipientID, amount)
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
