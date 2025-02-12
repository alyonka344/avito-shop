package repository

import (
	"avito-shop/internal/model"
	"github.com/gofrs/uuid/v5"
)

type UserRepository interface {
	Create(user *model.User) error
	Update(user *model.User) error
	GetById(userID uuid.UUID) (model.User, error)
	GetByName(userName string) (model.User, error)
}
