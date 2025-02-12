package repository

import "avito-shop/internal/model"

type merchRepository interface {
	GetByName(merchName string) (model.MerchItem, error)
}
