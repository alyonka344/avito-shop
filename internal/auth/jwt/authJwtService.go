package jwt

import (
	"avito-shop/internal/model"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"time"
)

const BearerString = "Bearer "

type JwtService struct {
	secretKey string
}

func NewJWTService(secretKey string) *JwtService {
	return &JwtService{secretKey: secretKey}
}

func (j *JwtService) GenerateToken(user model.User) (string, error) {
	claims := jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j *JwtService) ValidateToken(tokenStr string) (string, error) {
	tokenStr = strings.TrimPrefix(tokenStr, BearerString)

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(j.secretKey), nil
	})

	if err != nil || !token.Valid {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid token claims")
	}

	userName, ok := claims["username"].(string)
	if !ok {
		return "", errors.New("invalid username")
	}

	return userName, nil
}
