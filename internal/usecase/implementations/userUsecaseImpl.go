package implementations

import (
	"avito-shop/internal/repository"
	"avito-shop/internal/usecase"
)

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepository repository.UserRepository) usecase.UserUsecase {
	return &userUsecase{userRepository: userRepository}
}

//func (u *userUsecase) Create(user *model.User) error {
//	return u.userRepository.Create(user)
//}
