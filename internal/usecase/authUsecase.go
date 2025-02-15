package usecase

import (
	"avito-shop/internal/model"
)

type AuthUsecase interface {
	Login(email, password string) (string, error)
	ValidateOrCreateUser(username, password string) (model.User, error)
}
