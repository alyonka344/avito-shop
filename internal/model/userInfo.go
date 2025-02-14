package model

type UserInfo struct {
	Coins       int                 `json:"coins"`
	Inventory   []InventoryItem     `json:"inventory"`
	CoinHistory CoinHistoryResponse `json:"coinHistory"`
}

type InventoryItem struct {
	Type     string `json:"type"`
	Quantity int    `json:"quantity"`
}

type CoinHistoryResponse struct {
	Received []CoinTransaction `json:"received"`
	Sent     []CoinTransaction `json:"sent"`
}

type CoinTransaction struct {
	User   string `json:"user"`
	Amount int    `json:"amount"`
}
