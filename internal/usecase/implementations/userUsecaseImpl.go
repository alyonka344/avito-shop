package implementations

import (
	"avito-shop/internal/repository"
	"avito-shop/internal/usecase"
	"github.com/gofrs/uuid/v5"
)

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepository repository.UserRepository) usecase.UserUsecase {
	return &userUsecase{userRepository: userRepository}
}

func (u userUsecase) TransferMoney(senderID uuid.UUID, recipientID uuid.UUID, amount int) error {
	//if amount <= 0 {
	//	return errors.New("amount must be greater than zero")
	//}
	//
	//transaction := model.Transaction{
	//	ID:         uuid.Must(uuid.NewV4()),
	//	FromUserID: senderID,
	//	ToUserID:   recipientID,
	//	Amount:     amount,
	//	CreatedAt:  time.Now(),
	//}
	//
	//err := u.userRepository.Transfer(senderID, recipientID, amount)
	//if err != nil {
	//	transaction.TransactionStatus = model.Failure
	//	err = transactionRepo.LogTransaction(&transaction)
	//	if err != nil {
	//		return err
	//	}
	//	return err
	//}
	//
	//err = uc.transactionRepo.LogTransaction(&transaction)
	//if err != nil {
	//	return err
	//}

	return nil
}

//func (u *userUsecase) Create(user *model.User) error {
//	return u.userRepository.Create(user)
//}
