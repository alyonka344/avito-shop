package jwt

import (
	"avito-shop/internal/model"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofrs/uuid/v5"
	"time"
)

type JwtService struct {
	secretKey string
}

func NewJWTService(secretKey string) *JwtService {
	return &JwtService{secretKey: secretKey}
}

func (j *JwtService) GenerateToken(user model.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID.String(),
		"name":    user.Username,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j *JwtService) ValidateToken(tokenStr string) (uuid.UUID, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(j.secretKey), nil
	})

	if err != nil || !token.Valid {
		return uuid.Nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return uuid.Nil, errors.New("invalid token claims")
	}

	userIDStr, ok := claims["user_id"].(string)
	if !ok {
		return uuid.Nil, errors.New("invalid user_id")
	}

	userID, err := uuid.FromString(userIDStr)
	if err != nil {
		return uuid.Nil, errors.New("invalid user_id")
	}

	return userID, nil
}
