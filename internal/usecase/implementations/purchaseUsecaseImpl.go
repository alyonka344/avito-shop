package implementations

import (
	"avito-shop/internal/model"
	"avito-shop/internal/repository"
	"errors"
	"fmt"
)

type PurchaseUsecase struct {
	userRepository     repository.UserRepository
	merchRepository    repository.MerchRepository
	purchaseRepository repository.PurchaseRepository
}

func NewPurchaseUsecase(
	userRepository repository.UserRepository,
	merchRepository repository.MerchRepository,
	purchaseRepository repository.PurchaseRepository) *PurchaseUsecase {
	return &PurchaseUsecase{
		userRepository:     userRepository,
		merchRepository:    merchRepository,
		purchaseRepository: purchaseRepository,
	}
}

func (p PurchaseUsecase) BuyMerch(userName string, merchName string) error {
	merch, err := p.merchRepository.GetByName(merchName)
	if err != nil {
		fmt.Println("merch")
		return err
	}

	user, err := p.userRepository.GetByName(userName)
	if err != nil {
		fmt.Println("user")
		return err
	}

	if user.Balance < merch.Price {
		fmt.Println("balance")
		fmt.Println(user.Balance, merch.Price)
		return errors.New("insufficient balance")
	}

	err = p.userRepository.UpdateBalance(userName, -merch.Price)
	if err != nil {
		fmt.Println("update")
		return err
	}

	purchase := model.Purchase{
		UserName:  userName,
		MerchName: merchName,
	}
	return p.purchaseRepository.Create(&purchase)
}
