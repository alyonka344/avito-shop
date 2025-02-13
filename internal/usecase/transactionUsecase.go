package usecase

import "github.com/gofrs/uuid/v5"

type TransactionUsecase interface {
	TransferMoney(senderID uuid.UUID, recipientID uuid.UUID, amount int) error
}
