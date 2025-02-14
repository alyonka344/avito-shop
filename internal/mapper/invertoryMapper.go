package mapper

import "avito-shop/internal/model"

func MapInventory(purchases []model.Purchase) []model.InventoryItem {
	inventoryMap := make(map[string]int)
	for _, purchase := range purchases {
		inventoryMap[purchase.MerchName]++
	}

	var inventory []model.InventoryItem
	for itemType, quantity := range inventoryMap {
		inventory = append(inventory, model.InventoryItem{
			Type:     itemType,
			Quantity: quantity,
		})
	}
	return inventory
}
