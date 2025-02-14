package usecase

import (
	"avito-shop/internal/model"
)

type UserUsecase interface {
	GetInfo(userName string) (*model.UserInfo, error)
}
