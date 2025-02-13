package implementations

import (
	"avito-shop/internal/model"
	"avito-shop/internal/repository"
	"avito-shop/internal/usecase"
	"errors"
	"github.com/gofrs/uuid/v5"
)

type purchaseUsecase struct {
	userRepository     repository.UserRepository
	merchRepository    repository.MerchRepository
	purchaseRepository repository.PurchaseRepository
}

func NewPurchaseUsecase(
	userRepository repository.UserRepository,
	merchRepository repository.MerchRepository,
	purchaseRepository repository.PurchaseRepository) usecase.PurchaseUsecase {
	return &purchaseUsecase{
		userRepository:     userRepository,
		merchRepository:    merchRepository,
		purchaseRepository: purchaseRepository,
	}
}

func (p purchaseUsecase) BuyMerch(userID uuid.UUID, merchName string) error {
	merch, err := p.merchRepository.GetByName(merchName)
	if err != nil {
		return err
	}

	user, err := p.userRepository.GetById(userID)
	if err != nil {
		return err
	}

	if user.Balance < merch.Price {
		return errors.New("insufficient balance")
	}

	err = p.userRepository.UpdateBalance(userID, -merch.Price)
	if err != nil {
		return err
	}

	purchase := model.Purchase{
		UserID:    userID,
		MerchName: merchName,
	}
	return p.purchaseRepository.Create(&purchase)
}
