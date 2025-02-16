package unit_tests

import (
	"avito-shop/internal/mocks"
	"avito-shop/internal/model"
	"avito-shop/internal/usecase/usecase_impl"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTransactionUsecase_TransferMoneySuccessful(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mocks.NewMockUserRepository(ctrl)
	mockTransactionRepo := mocks.NewMockTransactionRepository(ctrl)

	usecase := usecase_impl.NewTransactionUseCase(mockUserRepo, mockTransactionRepo)

	senderName := "sender"
	recipientName := "recipient"
	amount := 100

	mockUserRepo.EXPECT().Transfer(senderName, recipientName, amount).Return(nil)

	// act
	mockTransactionRepo.EXPECT().
		Create(gomock.Any()).
		Do(func(transaction *model.Transaction) {
			assert.Equal(t, senderName, transaction.FromUser)
			assert.Equal(t, recipientName, transaction.ToUser)
			assert.Equal(t, amount, transaction.Amount)
			assert.Equal(t, model.Success, transaction.TransactionStatus)
		}).
		Return(nil)

	err := usecase.TransferMoney(senderName, recipientName, amount)

	// assert
	assert.NoError(t, err)
}

func TestTransactionUsecase_TestTransferMoneyNegativeAmount(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mocks.NewMockUserRepository(ctrl)
	mockTransactionRepo := mocks.NewMockTransactionRepository(ctrl)

	usecase := usecase_impl.NewTransactionUseCase(mockUserRepo, mockTransactionRepo)

	senderName := "sender"
	recipientName := "recipient"
	amount := -100

	// act
	err := usecase.TransferMoney(senderName, recipientName, amount)

	// assert
	assert.Error(t, err)
	assert.Equal(t, "amount must be greater than zero", err.Error())
}

func TestTransactionUsecase_TransferMoneyInsufficientFunds(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mocks.NewMockUserRepository(ctrl)
	mockTransactionRepo := mocks.NewMockTransactionRepository(ctrl)

	sender := &model.User{Username: "testUser1", Balance: 50}
	receiver := &model.User{Username: "testUser2", Balance: 100}

	mockUserRepo.EXPECT().Transfer(sender.Username, receiver.Username, 100).
		Return(errors.New("insufficient funds"))

	mockTransactionRepo.EXPECT().
		Create(gomock.Any()).
		Do(func(transaction *model.Transaction) {
			assert.Equal(t, sender.Username, transaction.FromUser)
			assert.Equal(t, receiver.Username, transaction.ToUser)
			assert.Equal(t, 100, transaction.Amount)
			assert.Equal(t, model.Failure, transaction.TransactionStatus)
		}).
		Return(nil)

	transactionUC := usecase_impl.NewTransactionUseCase(mockUserRepo, mockTransactionRepo)

	// Act
	err := transactionUC.TransferMoney(sender.Username, receiver.Username, 100)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, "transaction failed", err.Error())
}
