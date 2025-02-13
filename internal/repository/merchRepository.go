package repository

import "avito-shop/internal/model"

type MerchRepository interface {
	GetByName(merchName string) (*model.MerchItem, error)
}
