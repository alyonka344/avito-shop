package jwt

import (
	"avito-shop/internal/model"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strings"
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
		"username": user.Username,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j *JwtService) ValidateToken(tokenStr string) (string, error) {
	if strings.HasPrefix(tokenStr, "Bearer ") {
		tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")
	}

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("a")
			return nil, errors.New("unexpected signing method")
		}
		return []byte(j.secretKey), nil
	})

	if err != nil || !token.Valid {
		fmt.Println(err)
		fmt.Println(!token.Valid)
		fmt.Println("b")
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("c")
		return "", errors.New("invalid token claims")
	}

	userName, ok := claims["username"].(string)
	if !ok {
		fmt.Println("d")
		return "", errors.New("invalid username")
	}

	return userName, nil
}
