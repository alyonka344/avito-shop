package implementations

import (
	"avito-shop/internal/mapper"
	"avito-shop/internal/model"
	"avito-shop/internal/repository"
	"fmt"
)

type UserUsecase struct {
	userRepository        repository.UserRepository
	transactionRepository repository.TransactionRepository
	purchaseRepository    repository.PurchaseRepository
}

func NewUserUsecase(
	userRepository repository.UserRepository,
	transactionRepository repository.TransactionRepository,
	purhaseRepository repository.PurchaseRepository) *UserUsecase {
	return &UserUsecase{
		userRepository:        userRepository,
		transactionRepository: transactionRepository,
		purchaseRepository:    purhaseRepository}
}

func (u UserUsecase) GetInfo(userName string) (*model.UserInfo, error) {
	user, err := u.userRepository.GetByName(userName)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	sentTransactions, err := u.transactionRepository.GetAllSentByUserName(userName)
	if err != nil {
		return nil, fmt.Errorf("failed to get sentTransactions: %w", err)
	}

	receivedTransactions, err := u.transactionRepository.GetAllReceivedByUserName(userName)
	if err != nil {
		return nil, fmt.Errorf("failed to get receivedTransactions: %w", err)
	}

	inventory, err := u.purchaseRepository.GetAllByUserName(userName)
	if err != nil {
		return nil, fmt.Errorf("failed to get inventory: %w", err)
	}

	userInfo := &model.UserInfo{
		Coins:     user.Balance,
		Inventory: mapper.MapInventory(inventory),
		CoinHistory: model.CoinHistoryResponse{
			Sent:     mapper.MapCoinTransactions(sentTransactions, true),
			Received: mapper.MapCoinTransactions(receivedTransactions, false),
		},
	}

	return userInfo, nil
}
