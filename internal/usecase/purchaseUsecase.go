package usecase

type PurchaseUsecase interface {
	BuyMerch(userName string, merchName string) error
}
