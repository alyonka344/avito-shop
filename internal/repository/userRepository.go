package repository

import (
	"avito-shop/internal/model"
	"github.com/gofrs/uuid/v5"
)

type UserRepository interface {
	Create(user *model.User) error
	GetById(userID uuid.UUID) (*model.User, error)
	GetByName(userName string) (*model.User, error)
	Transfer(senderID uuid.UUID, recipientID uuid.UUID, amount int) error
	UpdateBalance(userID uuid.UUID, amount int) error
	ExistsByName(userName string) (bool, error)
}
