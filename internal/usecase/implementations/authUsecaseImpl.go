package implementations

import (
	"avito-shop/internal/auth"
	"avito-shop/internal/model"
	"avito-shop/internal/repository"
	"avito-shop/internal/usecase"
	"errors"
)

type authUsecase struct {
	userRepository repository.UserRepository
	authService    auth.AuthService
	hashService    auth.HashService
}

func NewAuthUsecase(
	userRepository repository.UserRepository,
	authService auth.AuthService,
	hashService auth.HashService) usecase.AuthUsecase {
	return &authUsecase{userRepository: userRepository, authService: authService, hashService: hashService}
}

func (u *authUsecase) Register(user *model.User) error {
	hashedPassword, err := u.hashService.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return u.userRepository.Create(user)
}

func (u *authUsecase) Login(name, password string) (string, error) {
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

func (u *authUsecase) ValidateOrCreateUser(username, password string) (model.User, error) {
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

		return *newUser, err
	}
}
