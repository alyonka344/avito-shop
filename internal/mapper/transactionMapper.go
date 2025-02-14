package mapper

import (
	"avito-shop/internal/model"
)

func MapCoinTransactions(transactions []model.Transaction, isSent bool) []model.CoinTransaction {
	var coinTransactions []model.CoinTransaction
	for _, tx := range transactions {
		var userName string
		if isSent {
			userName = tx.ToUser
		} else {
			userName = tx.FromUser
		}

		coinTransactions = append(coinTransactions, model.CoinTransaction{
			User:   userName,
			Amount: tx.Amount,
		})
	}
	return coinTransactions
}
