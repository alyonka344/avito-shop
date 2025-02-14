package usecase

type TransactionUsecase interface {
	TransferMoney(senderName string, recipientName string, amount int) error
}
