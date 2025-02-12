package implementations

import (
	"avito-shop/internal/auth"
	"avito-shop/internal/model"
	"avito-shop/internal/repository"
	"avito-shop/internal/usecase"
	"errors"
)

type authUseCase struct {
	userRepository repository.UserRepository
	authService    auth.AuthService
	hashService    auth.HashService
}

func NewAuthUsecase(
	userRepository repository.UserRepository,
	authService auth.AuthService,
	hashService auth.HashService) usecase.AuthUseCase {
	return &authUseCase{userRepository: userRepository, authService: authService, hashService: hashService}
}

func (u *authUseCase) Register(user *model.User) error {
	hashedPassword, err := u.hashService.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return u.userRepository.Create(user)
}

func (u *authUseCase) Login(name, password string) (string, error) {
	user, err := u.userRepository.GetByName(name)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if err := u.hashService.CompareHashAndPassword(user.Password, password); err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := u.authService.GenerateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}
