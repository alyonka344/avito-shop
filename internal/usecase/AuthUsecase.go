package usecase

import "avito-shop/internal/model"

type AuthUseCase interface {
	Login(email, password string) (string, error)
	Register(user *model.User) error
}
