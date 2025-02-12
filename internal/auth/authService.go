package auth

import (
	"avito-shop/internal/model"
	"github.com/gofrs/uuid/v5"
)

type AuthService interface {
	GenerateToken(user model.User) (string, error)
	ValidateToken(token string) (uuid.UUID, error)
}
