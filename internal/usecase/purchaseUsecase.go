package usecase

import "github.com/gofrs/uuid/v5"

type PurchaseUsecase interface {
	BuyMerch(userID uuid.UUID, merchName string) error
}
