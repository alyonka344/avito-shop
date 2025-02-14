package implementations

import (
	"avito-shop/internal/model"
	"avito-shop/internal/repository"
	"avito-shop/internal/usecase"
	"errors"
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

func (p purchaseUsecase) BuyMerch(userName string, merchName string) error {
	merch, err := p.merchRepository.GetByName(merchName)
	if err != nil {
		return err
	}

	user, err := p.userRepository.GetByName(userName)
	if err != nil {
		return err
	}

	if user.Balance < merch.Price {
		return errors.New("insufficient balance")
	}

	err = p.userRepository.UpdateBalance(userName, -merch.Price)
	if err != nil {
		return err
	}

	purchase := model.Purchase{
		UserName:  userName,
		MerchName: merchName,
	}
	return p.purchaseRepository.Create(&purchase)
}
