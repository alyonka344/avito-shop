package auth

import (
	"avito-shop/internal/model"
)

type AuthService interface {
	GenerateToken(user model.User) (string, error)
	ValidateToken(token string) (string, error)
}
