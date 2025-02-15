package implementations

import (
	"avito-shop/internal/auth"
	"avito-shop/internal/model"
	"avito-shop/internal/repository"
	"errors"
	"fmt"
)

type AuthUsecase struct {
	userRepository repository.UserRepository
	authService    auth.AuthService
	hashService    auth.HashService
}

func NewAuthUsecase(
	userRepository repository.UserRepository,
	authService auth.AuthService,
	hashService auth.HashService) *AuthUsecase {
	return &AuthUsecase{
		userRepository: userRepository,
		authService:    authService,
		hashService:    hashService}
}

func (u *AuthUsecase) Login(name, password string) (string, error) {
	user, err := u.userRepository.GetByName(name)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if err := u.hashService.CompareHashAndPassword(user.Password, password); err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := u.authService.GenerateToken(*user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *AuthUsecase) ValidateOrCreateUser(username, password string) (model.User, error) {
	exists, err := u.userRepository.ExistsByName(username)
	if err != nil {
		return model.User{}, err
	}
	if exists {
		user, err := u.userRepository.GetByName(username)
		if err == nil {
			return *user, nil
		}
		return model.User{}, err
	} else {

		hashedPassword, err := u.hashService.HashPassword(password)
		if err != nil {
			return model.User{}, err
		}

		newUser := &model.User{
			Username: username,
			Password: hashedPassword,
		}

		if err := u.userRepository.Create(newUser); err != nil {
			return model.User{}, err
		}
		fmt.Println(newUser.Balance)

		return *newUser, err
	}
}
