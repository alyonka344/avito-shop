package repository

import (
	"avito-shop/internal/model"
)

type UserRepository interface {
	Create(user *model.User) error
	GetByName(userName string) (*model.User, error)
	Transfer(senderName string, recipientName string, amount int) error
	UpdateBalance(userName string, amount int) error
	ExistsByName(userName string) (bool, error)
}
